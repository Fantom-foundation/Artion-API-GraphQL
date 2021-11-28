// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"time"
)

// zeroAddress represents an empty address.
var zeroAddress = common.Address{}

// erc721TokenMinted handles log event for new NFT token minted on an observed ERC721 contract.
// ERC721::Minted(uint256 tokenId, address beneficiary, string tokenUri, address minter)
func erc721TokenMinted(evt *eth.Log, lo *logObserver) {
	// sanity check: no extra topics; tokenId + 2 x Address + URI >= 3 x 32 bytes
	if len(evt.Data) < 96 || len(evt.Topics) != 1 {
		log.Errorf("not ERC721::Minted() event #%d/#%d; expected at least 96 bytes of data, %d given; expected 1 topic, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	// unpack the event data
	args, err := repository.R().Erc721Abi().Unpack("Minted", evt.Data)
	if err != nil {
		log.Errorf("can not decode ERC721 %s mint data; %s", evt.Address.String(), err.Error())
		return
	}

	// get the block header
	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("can not load event header #%d; %s", evt.BlockNumber, err.Error())
		return
	}

	// make the token
	tok := types.NewToken(&evt.Address, args[0].(*big.Int), args[2].(string), int64(blk.Time), evt.BlockNumber, evt.Index)
	log.Infof("ERC-721 token %s found at %s block %d", tok.TokenId.String(), tok.Contract.String(), evt.BlockNumber)

	// store the extra data
	tok.CreatedBy = args[3].(common.Address)

	if err := repo.TokenLikesViewsRefresh(tok); err != nil {
		log.Errorf("could not load token views/likes %s/%s; %s", tok.TokenId.String(), tok.Contract.String(), err)
	}

	// write token to the persistent storage
	if err := repo.StoreToken(tok); err != nil {
		log.Errorf("could not store token %s at %s; %s", tok.TokenId.String(), tok.Contract.String(), err.Error())
		return
	}

	// schedule metadata update on the token (do not wait for result)
	queueMetadataUpdate(tok, lo)
}

// erc721TokenMustExist ensures ERC-721 token existence in the local database.
func erc721TokenMustExist(contract *common.Address, tokenID *big.Int, blk *eth.Header, evt *eth.Log, lo *logObserver) {
	log.Debugf("checking %s / #%d", contract.String(), tokenID.Uint64())
	if repo.TokenKnown(contract, tokenID) {
		log.Debugf("found token %s / #%d", contract.String(), tokenID.Uint64())
		return
	}

	// get token URI
	uri, err := repo.Erc721TokenUri(contract, tokenID)
	if err != nil {
		log.Errorf("could not get token URI for %s / #%d; %s", contract.String(), tokenID.Uint64(), err.Error())
		uri = ""
	}

	// make the token
	tok := types.NewToken(contract, tokenID, uri, int64(blk.Time), evt.BlockNumber, evt.Index)
	log.Infof("ERC-721 token %s found at %s block %d", tok.TokenId.String(), tok.Contract.String(), evt.BlockNumber)

	// add details
	tok.CreatedBy = repo.MustTransactionSender(evt.TxHash)

	if err := repo.TokenLikesViewsRefresh(tok); err != nil {
		log.Errorf("could not load token views/likes %s/%s; %s", tok.TokenId.String(), tok.Contract.String(), err)
	}

	// write token to the persistent storage
	if err := repo.StoreToken(tok); err != nil {
		log.Errorf("could not store token %s at %s; %s", tok.TokenId.String(), tok.Contract.String(), err.Error())
		return
	}

	// schedule metadata update on the token (do not wait for result)
	queueMetadataUpdate(tok, lo)
}

// erc721TokenTransfer handles log event for NFT token ownership transfer on an observed ERC721 contract.
// ERC721::Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func erc721TokenTransfer(evt *eth.Log, lo *logObserver) {
	// sanity check: 1 + 3 extra topics for indexed parties; no additional data = 0 bytes
	if len(evt.Data) != 0 || len(evt.Topics) != 4 {
		log.Errorf("not ERC721::Transfer() event #%d/#%d; expected no data, %d given; expected 4 topics, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	// extract details
	from := common.BytesToAddress(evt.Topics[1].Bytes())
	to := common.BytesToAddress(evt.Topics[2].Bytes())
	tokenID := hexutil.Big(*new(big.Int).SetBytes(evt.Topics[3].Bytes()))

	// get the block header
	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("can not load event header #%d; %s", evt.BlockNumber, err.Error())
		return
	}

	// log what we do
	log.Debugf("erc721 %s / %s transfer %s -> %s", evt.Address.String(), tokenID.String(), from.String(), to.String())

	// this may be a mint
	// it could have been covered on artion NFT contracts above, but we need to handle non-artion contracts, too
	if bytes.Equal(zeroAddress.Bytes(), from.Bytes()) {
		erc721TokenMustExist(&evt.Address, (*big.Int)(&tokenID), blk, evt, lo)
		notifyEventToOwner(types.NotifyNFTCreated, evt.Address, tokenID, to, nil, types.Time(time.Unix(int64(blk.Time), 0)))
	}

	// ERC-721 tokens don't have quantity; the amount is always 1
	// we can just clear previous owner here by setting qty to zero
	if err := updateERC721Owner(evt.Address, tokenID, from, 0, blk.Time); err != nil {
		log.Errorf("could not clear ERC-721 NFT ownership; %s", err.Error())
		return
	}

	// is this a token burn event?
	if bytes.Equal(zeroAddress.Bytes(), to.Bytes()) {
		if err := registerERC721TokenBurn(evt.Address, tokenID, from, blk.Time); err != nil {
			log.Errorf("could not add ERC-721 NFT burn; %s", err.Error())
		}
	} else {
		// now we can add the new owner
		if err := updateERC721Owner(evt.Address, tokenID, to, 1, blk.Time); err != nil {
			log.Errorf("could not add ERC-721 NFT ownership; %s", err.Error())
			return
		}
	}

	// update listings/auctions status
	setAuctionsListingsActive(&evt.Address, (*big.Int)(&tokenID), &from, &to)

	// log transfer activity
	qty := big.NewInt(1) // 1 for every ERC-721 transfer
	logTokenTransferActivity(evt, blk, tokenID, qty, from, to)
}

// queueMetadataUpdate pushes NFT into the metadata processing queue.
func queueMetadataUpdate(nft *types.Token, lo *logObserver) {
	// schedule metadata update on the token (do not wait for result)
	// if the updater queue is full, we just let the updater pick the token for update later
	select {
	case lo.outNftTokens <- nft:
	default:
		log.Errorf("NFT token updater queue full, postponing token %s at %s metadata update", nft.TokenId.String(), nft.Contract.String())
	}
}

// updateERC721Owner pushes the ownership of ERC-721 token into persistent storage.
func updateERC721Owner(contract common.Address, tokenID hexutil.Big, owner common.Address, qty uint64, ts uint64) error {
	// now we can add the new owner
	if err := repo.StoreOwnership(&types.Ownership{
		Contract: contract,
		TokenId:  tokenID,
		Owner:    owner,
		Qty:      hexutil.Big(*new(big.Int).SetUint64(qty)),
		Updated:  types.Time(time.Unix(int64(ts), 0)),
	}); err != nil {
		return err
	}
	return nil
}

// registerERC721TokenBurn registers a burn event on NFT.
func registerERC721TokenBurn(contract common.Address, tokenID hexutil.Big, owner common.Address, ts uint64) error {
	// store the burn
	if err := repo.StoreBurn(&types.NFTBurn{
		Contract: contract,
		TokenId:  tokenID,
		Owner:    owner,
		Qty:      hexutil.Big(*new(big.Int).SetUint64(1)),
		Burned:   types.Time(time.Unix(int64(ts), 0)),
	}); err != nil {
		return err
	}

	// notify burn via notifications to the owner
	notifyEventToOwner(types.NotifyNFTBurned, contract, tokenID, owner, nil, types.Time(time.Unix(int64(ts), 0)))
	return nil
}

// logTokenTransferActivity inserts transfer/mint/burn record into activities
func logTokenTransferActivity(evt *eth.Log, blk *eth.Header, tokenID hexutil.Big, qty *big.Int, from common.Address, to common.Address) {
	actType := types.EvtTransfer
	if bytes.Equal(zeroAddress.Bytes(), from.Bytes()) {
		actType = types.EvtMint
	}
	if bytes.Equal(zeroAddress.Bytes(), to.Bytes()) {
		actType = types.EvtBurn
	}
	activity := types.Activity{
		OrdinalIndex: types.OrdinalIndex(int64(evt.BlockNumber), int64(evt.Index)),
		Time:         types.Time(time.Unix(int64(blk.Time), 0)),
		ActType:      actType,
		Contract:     evt.Address,
		TokenId:      tokenID,
		Quantity:     (*hexutil.Big)(qty),
		From:         from,
		To:           &to,
	}
	if err := repo.StoreActivity(&activity); err != nil {
		log.Errorf("could not store transfer activity; %s", err.Error())
		return
	}
}

// setAuctionsListingsActive sets listings/auction inactive for not-already-owner (and back)
func setAuctionsListingsActive(contract *common.Address, tokenID *big.Int, from *common.Address, to *common.Address) {
	if err := repository.R().SetListingActive(contract, tokenID, from, false); err != nil {
		log.Errorf("unable to update listing active status on ownership change; %s", err.Error())
	}
	if err := repository.R().SetListingActive(contract, tokenID, to, true); err != nil {
		log.Errorf("unable to update listing active status on ownership change; %s", err.Error())
	}
	if err := repository.R().SetAuctionActive(contract, tokenID, from, false); err != nil {
		log.Errorf("unable to update auction active status on ownership change; %s", err.Error())
	}
	if err := repository.R().SetAuctionActive(contract, tokenID, to, true); err != nil {
		log.Errorf("unable to update auction active status on ownership change; %s", err.Error())
	}

	// update the token price after listings/auction changes
	if err := repo.TokenMarkSold(
		contract,
		tokenID,
		nil,
		nil,
	); err != nil {
		log.Errorf("could not update token price; %s", err.Error())
	}
}
