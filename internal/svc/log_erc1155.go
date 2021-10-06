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
	from := new(common.Address)
	from.SetBytes(evt.Topics[2].Bytes())
	to := new(common.Address)
	to.SetBytes(evt.Topics[3].Bytes())
	tokenId := new(big.Int).SetBytes(evt.Data[:32])

	// add recipient ownership record; we can do it first even for new tokens
	if err := repo.StoreOwnership(&types.Ownership{
		Contract: evt.Address,
		TokenId:  hexutil.Big(*tokenId),
		Owner:    *to,
		Qty:      repo.BalanceAt(&evt.Address, tokenId, to, &evt.BlockNumber),
		Updated:  types.Time(time.Now()),
	}); err != nil {
		log.Errorf("failed to update recipient ownership at %s/%d; %s",
			evt.Address.String(), tokenId.Uint64(), err.Error())
	}

	// make sure we know the token
	// if this is a mint call, the sender is zero, and we just add the NFT record
	if 0 == bytes.Compare(zeroAddress.Bytes(), from.Bytes()) {
		addERC1155Token(&evt.Address, tokenId)
		return
	}

	// if it's not a mint, we need to update sender's ownership balance as well
	if err := repo.StoreOwnership(&types.Ownership{
		Contract: evt.Address,
		TokenId:  hexutil.Big(*tokenId),
		Owner:    *to,
		Qty:      repo.BalanceAt(&evt.Address, tokenId, from, &evt.BlockNumber),
		Updated:  types.Time(time.Now()),
	}); err != nil {
		log.Errorf("failed to update sender ownership at %s/%d; %s",
			evt.Address.String(), tokenId.Uint64(), err.Error())
	}
}

// addERC1155Token adds a new ERC1155 type of token into the repository.
func addERC1155Token(adr *common.Address, tokenID *big.Int) {
	// make the token
	tok := types.Token{
		Contract: *adr,
		TokenId:  hexutil.Big(*tokenID),
		Uri:      repo.TokenUri(adr, tokenID),
	}

	tok.GenerateId()
	log.Infof("ERC-1155 token #%s found at %s", tok.TokenId.String(), tok.Contract.String())

	// write token to the persistent storage
	if err := repo.StoreToken(&tok); err != nil {
		log.Errorf("could not store token %s at %s; %s", tok.TokenId.String(), tok.Contract.String(), err.Error())
	}
}
