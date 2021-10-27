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
var zeroAddress common.Address

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
func erc721TokenTransfer(evt *eth.Log, _ *logObserver) {
	// sanity check: 1 + 3 extra topics for indexed parties; no additional data = 0 bytes
	if len(evt.Data) != 0 || len(evt.Topics) != 4 {
		log.Errorf("not ERC721::Transfer() event #%d/#%d; expected no data, %d given; expected 4 topics, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	// this may be a mint; if so, we have that one covered already
	if 0 == bytes.Compare(zeroAddress.Bytes(), evt.Topics[1].Bytes()) {
		log.Debug("ERC721::Mint() detected by token transfer")
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

	// ERC-721 tokens don't have quantity; the amount is always 1
	// we can just clear previous owner here by setting qty to zero
	if err := updateERC721Owner(evt.Address, tokenID, from, 0, blk.Time); err != nil {
		log.Errorf("could not clear ERC-721 NFT ownership; %s", err.Error())
		return
	}

	// is this a token burn event?
	if 0 == bytes.Compare(to.Bytes(), zeroAddress.Bytes()) {
		if err := registerERC721TokenBurn(evt.Address, tokenID, from, blk.Time); err != nil {
			log.Errorf("could not add ERC-721 NFT burn; %s", err.Error())
		}
		return
	}

	// now we can add the new owner
	if err := updateERC721Owner(evt.Address, tokenID, to, 1, blk.Time); err != nil {
		log.Errorf("could not add ERC-721 NFT ownership; %s", err.Error())
		return
	}
}

// queueMetadataUpdate pushes NFT into the metadata processing queue.
func queueMetadataUpdate(nft *types.Token, lo *logObserver) {
	// schedule metadata update on the token (do not wait for result)
	// if the updater queue is full, we just let the updater pick the token for update later
	select {
	case lo.outNftTokens <- nft:
	case <-time.After(10 * time.Second):
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

// registerTokenBurn registers
func registerERC721TokenBurn(contract common.Address, tokenID hexutil.Big, owner common.Address, ts uint64) error {
	// now we can add the new owner
	if err := repo.StoreBurn(&types.NFTBurn{
		Contract: contract,
		TokenId:  tokenID,
		Owner:    owner,
		Qty:      hexutil.Big(*new(big.Int).SetUint64(1)),
		Burned:   types.Time(time.Unix(int64(ts), 0)),
	}); err != nil {
		return err
	}

	// notify burn via notifications
	repo.QueueNotificationForProcessing(&types.Notification{
		Type:       types.NotifyBurnedNFTToken,
		Contract:   contract,
		TokenId:    tokenID,
		TimeStamp:  types.Time(time.Unix(int64(ts), 0)),
		Recipient:  owner,
		Originator: nil,
	})
	return nil
}
