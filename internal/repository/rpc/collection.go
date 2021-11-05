// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/status-im/keycard-go/hexutils"
)

// CollectionName provides a name of an Artion ERC721 and/or ERC1155 token.
// Solidity: function name() view returns(string)
func (o *Opera) CollectionName(adr *common.Address) (string, error) {
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   adr,
		Data: hexutils.HexToBytes("06fdde03"),
	}, nil)
	if err != nil {
		log.Errorf("contract %s name not found", adr.String())
		return "", err
	}
	res, err := o.abiFantom721.Unpack("name", data)
	if err != nil {
		log.Errorf("can not decode contract %s name; %s", adr.String(), err.Error())
		return "", err
	}
	return *abi.ConvertType(res[0], new(string)).(*string), nil
}

// CollectionSymbol provides a symbol of an Artion ERC721 and/or ERC1155 token.
// Solidity: function symbol() view returns(string)
func (o *Opera) CollectionSymbol(adr *common.Address) (string, error) {
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   adr,
		Data: hexutils.HexToBytes("95d89b41"),
	}, nil)
	if err != nil {
		log.Errorf("contract %s symbol not found", adr.String())
		return "", err
	}
	res, err := o.abiFantom721.Unpack("symbol", data)
	if err != nil {
		log.Errorf("can not decode contract %s symbol; %s", adr.String(), err.Error())
		return "", err
	}
	return *abi.ConvertType(res[0], new(string)).(*string), nil
}

// SupportsInterface checks if the given address is a contract
// implementing a specified interface using ERC-165 contract call.
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (o *Opera) SupportsInterface(adr *common.Address, in string) bool {
	// we need 4 bytes for function hash (0x01ffc9a7) + 32 bytes for the interface hex input (only 4 are used)
	call := make([]byte, 32+4)
	copy(call[:4], []byte{0x01, 0xff, 0xc9, 0xa7})
	if in[0:2] == "0x" {
		in = in[2:]
	}
	copy(call[4:], hexutils.HexToBytes(in)[:4])

	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   adr,
		Data: call,
	}, nil)
	if err != nil {
		log.Warningf("erc-721 check failed; %s", err.Error())
		return false
	}

	// we expect 32 bytes 0x0000000000000000000000000000000000000000000000000000000000000001 for TRUE
	log.Infof("%s (%s) responded with %s", adr.String(), in, hexutils.BytesToHex(data))
	return len(data) == 32 && data[0] == 0 && data[31] > 0
}
