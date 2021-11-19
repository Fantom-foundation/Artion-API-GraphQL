// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/types"
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"time"
)

// erc1155TokenTransfer handles single ERC1155 NFT token transfer
// including a new token Mint(), if the sender is zero address.
// ERC1155::TransferSingle(address indexed _operator, address indexed _from, address indexed _to, uint256 _id, uint256 _amount)
func erc1155TokenTransfer(evt *eth.Log, lo *logObserver) {
	// sanity check: 1 + 3 topics; 2 x uint256 = 2 x 32 bytes of data
	if len(evt.Data) != 34 || len(evt.Topics) != 4 {
		log.Errorf("not ERC1155::TransferSingle() event #%d / #%d; expected 64 bytes of data, %d given; expected 4 topics, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	// extract data
	from := common.BytesToAddress(evt.Topics[2].Bytes())
	to := common.BytesToAddress(evt.Topics[3].Bytes())
	tokenId := new(big.Int).SetBytes(evt.Data[:32])
	amount := new(big.Int).SetBytes(evt.Data[32:64])

	// add recipient ownership record; we can do it first even for new tokens
	if err := repo.StoreOwnership(&types.Ownership{
		Contract: evt.Address,
		TokenId:  hexutil.Big(*tokenId),
		Owner:    to,
		Qty:      balanceOf(&evt.Address, tokenId, &to, evt.BlockNumber, lo),
		Updated:  types.Time(time.Now()),
	}); err != nil {
		log.Errorf("failed to update recipient ownership at %s/%d; %s",
			evt.Address.String(), tokenId.Uint64(), err.Error())
	}

	// make sure we know the token
	// if this is a mint call (the sender is zero), we just add the NFT record
	if bytes.Equal(zeroAddress.Bytes(), from.Bytes()) {
		addERC1155Token(&evt.Address, tokenId, &to, evt, lo)
	} else {
		// if it's not a mint, we need to update sender's ownership balance as well
		if err := repo.StoreOwnership(&types.Ownership{
			Contract: evt.Address,
			TokenId:  hexutil.Big(*tokenId),
			Owner:    to,
			Qty:      balanceOf(&evt.Address, tokenId, &from, evt.BlockNumber, lo),
			Updated:  types.Time(time.Now()),
		}); err != nil {
			log.Errorf("failed to update sender ownership at %s/%d; %s",
				evt.Address.String(), tokenId.Uint64(), err.Error())
		}
	}

	// get the block header
	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("can not load event header #%d; %s", evt.BlockNumber, err.Error())
		return
	}

	// log transfer activity
	logTokenTransferActivity(evt, blk, hexutil.Big(*tokenId), amount, from, to)
}

// balanceOf returns the balance of a token for the given owner on the given block.
func balanceOf(con *common.Address, tokenId *big.Int, owner *common.Address, block uint64, _ *logObserver) hexutil.Big {
	// try to get the contract type
	tt, err := repo.ContractTypeByAddress(con)
	if err != nil {
		log.Criticalf("unknown contract type; %s", err.Error())
		return hexutil.Big{}
	}

	switch tt {
	case types.ContractTypeERC721:
		// @todo If the owner has the token, return 1; if not, return 0
	case types.ContractTypeERC1155:
		qty, err := repo.Erc1155BalanceOf(con, tokenId, owner, new(big.Int).SetUint64(block))
		if err != nil {
			log.Criticalf("token balance unknown; %s", err.Error())
			return hexutil.Big{}
		}
		return hexutil.Big(*qty)
	default:
		log.Criticalf("unknown contract type %s", tt)
	}
	return hexutil.Big{}
}

// addERC1155Token adds a new ERC1155 type of token into the repository.
func addERC1155Token(adr *common.Address, tokenID *big.Int, creator *common.Address, evt *eth.Log, lo *logObserver) {
	// extract the token URI from the contract
	uri, err := repo.Erc1155TokenUri(adr, tokenID)
	if err != nil {
		log.Errorf("token URI not known; %s", err.Error())
	}

	// get the block header
	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("can not load event header #%d; %s", evt.BlockNumber, err.Error())
		return
	}

	// make the token
	tok := types.NewToken(adr, tokenID, uri, int64(blk.Time), evt.BlockNumber, evt.Index)
	tok.CreatedBy = *creator

	log.Infof("ERC-1155 token #%s found at %s", tok.TokenId.String(), tok.Contract.String())

	if err := repo.TokenLikesViewsRefresh(tok); err != nil {
		log.Errorf("could not load token views/likes %s/%s; %s", tok.TokenId.String(), tok.Contract.String(), err)
	}

	// write token to the persistent storage
	if err := repo.StoreToken(tok); err != nil {
		log.Errorf("could not store token %s at %s; %s", tok.TokenId.String(), tok.Contract.String(), err.Error())
	}

	// queue the token for metadata update
	queueMetadataUpdate(tok, lo)
}
