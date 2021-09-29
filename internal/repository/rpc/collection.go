// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/status-im/keycard-go/hexutils"
)

// CollectionName provides a name of an Artion ERC721 and/or ERC1155 token.
// Solidity: function name() view returns(string)
func (o *Opera) CollectionName(adr *common.Address) (string, error) {
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   adr,
		Data: hexutils.HexToBytes("0x06fdde03"),
	}, nil)
	if err != nil {
		return "", err
	}
	return string(data[:]), nil
}

// CollectionSymbol provides a symbol of an Artion ERC721 and/or ERC1155 token.
// Solidity: function symbol() view returns(string)
func (o *Opera) CollectionSymbol(adr *common.Address) (string, error) {
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   adr,
		Data: hexutils.HexToBytes("0x95d89b41"),
	}, nil)
	if err != nil {
		return "", err
	}
	return string(data[:]), nil
}
