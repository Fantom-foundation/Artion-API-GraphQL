// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// erc721TokenMinted handles log event for new NFT token minted on an observed ERC721 contract.
// ERC721::Minted(uint256 tokenId, address beneficiary, string tokenUri, address minter)
func erc721TokenMinted(evt *eth.Log, lo *logObserver) {
	// sanity check: no extra topics; tokenId + 2 x Address + URI >= 3 x 32 bytes
	if len(evt.Data) < 96 || len(evt.Topics) != 1 {
		log.Errorf("invalid event %s / %d; expected at least 96 bytes of data, %d given; expected 1 topic, %d given",
			evt.TxHash.String(), evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	// unpack the event data
	args, err := repository.R().Erc721Abi().Unpack("Minted", evt.Data)
	if err != nil {
		log.Errorf("can not decode ERC721 %s mint data; %s", evt.Address.String(), err.Error())
		return
	}

	// make the token
	tok := types.Token{
		Nft:     evt.Address,
		TokenId: hexutil.Big(*(args[0].(*big.Int))),
		Uri:     args[2].(string),
	}
	tok.GenerateId()
	log.Infof("ERC-721 token %s found at %s block %d", tok.TokenId.String(), tok.Nft.String(), evt.BlockNumber)

	// write token to the persistent storage
	if err := repo.StoreToken(&tok); err != nil {
		log.Errorf("could not store token %s at %s; %s", tok.TokenId.String(), tok.Nft.String(), err.Error())
		return
	}

	// schedule metadata update on the token (do not wait for result)
	// if the updater queue is full, we just let the updater pick the token for update later
	select {
	case lo.outNftTokens <- &tok:
	default:
		log.Errorf("NFT token updater queue full, postponing token %s at %s metadata update", tok.TokenId.String(), tok.Nft.String())
	}
}
