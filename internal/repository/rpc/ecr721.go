// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"artion-api-graphql/internal/repository/rpc/contracts"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// defaultMintingTestTokenUrl is the URL we used to test NFT minting calls.
const defaultMintingTestTokenUrl = "https://minter.artion.io/default/access/minter/estimation.json"

// Erc721StartingBlockNumber provides the first important block number for the ERC-721 contract.
// We try to get the first Transfer() event on the contract,
// anything before it is irrelevant for this API.
func (o *Opera) Erc721StartingBlockNumber(adr *common.Address) (uint64, error) {
	// instantiate contract
	erc, err := contracts.NewErc721(*adr, o.ftm)
	if err != nil {
		return 0, err
	}

	// iterate over transfers from zero address (e.g. mint calls)
	iter, err := erc.FilterTransfer(nil, []common.Address{{}}, nil, nil)
	if err != nil {
		return 0, err
	}

	var blk uint64
	if iter.Next() {
		blk = iter.Event.Raw.BlockNumber
	}

	if err := iter.Close(); err != nil {
		log.Errorf("could not close filter iterator; %s", err.Error())
	}
	return blk, nil
}

// CanMintErc721 checks if the given user can mint a new token on the given NFT contract.
func (o *Opera) CanMintErc721(contract *common.Address, user *common.Address) (bool, error) {
	data, err := o.abiFantom721.Pack("mint", *user, defaultMintingTestTokenUrl)
	if err != nil {
		return false, err
	}

	// try to estimate the call
	gas, err := o.ftm.EstimateGas(context.Background(), ethereum.CallMsg{
		From: *user,
		To:   contract,
		Data: data,
	})
	if err != nil {
		log.Warningf("user %s can not mint on ERC-721 %s; %s", user.String(), contract.String(), err.Error())
		return false, nil
	}

	log.Infof("user %s can mint on ERC-721 %s for %d gas", user.String(), contract.String(), gas)
	return true, nil
}
