// Package svc implements monitoring and scanning services of the API server.
package svc

import eth "github.com/ethereum/go-ethereum/core/types"

// newNFTContract handles log event for new factory deployed ERC721/ERC1155 contract.
// event ContractCreated(address creator, address nft)
func newNFTContract(evt *eth.Log) {

}