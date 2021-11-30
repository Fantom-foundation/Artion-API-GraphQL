// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/types"
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"time"
)

// IsObservedContract checks if the given address is a known and observed contract.
func (p *Proxy) IsObservedContract(adr *common.Address) bool {
	for _, con := range p.observedContracts {
		if 0 == bytes.Compare(adr.Bytes(), con.Bytes()) {
			return true
		}
	}
	return false
}

// AddObservedContract adds new observed contract into the persistent storage.
func (p *Proxy) AddObservedContract(oc *types.ObservedContract) {
	if p.IsObservedContract(&oc.Address) {
		return
	}

	err := p.db.AddObservedContract(oc)
	if err != nil {
		log.Criticalf("can not add observed contract %s; %s", oc.Address.String(), err.Error())
	}

	p.addObservedContract(&oc.Address, oc.Type)
	log.Infof("%s contract %s is now observed since #%d", oc.Type, oc.Address.String(), oc.BlockNumber)
}

// addObservedContract is used to extend the list of observed contracts
// with a newly created NFT contract address; subsequent NFT events should be observed on it.
func (p *Proxy) addObservedContract(adr *common.Address, ct string) {
	// check if the contract is actually a new one
	if p.IsObservedContract(adr) {
		return
	}

	// add the contract to the list
	p.observedContracts = append(p.observedContracts, *adr)

	// an NFT contract? add it to the types map as well
	if ct == types.ContractTypeERC721 || ct == types.ContractTypeERC1155 {
		p.nftTypes[*adr] = ct
	}
}

// ObservedContractsAddressList provides list of addresses of all observed contracts
// stored in the persistent storage.
func (p *Proxy) ObservedContractsAddressList() []common.Address {
	return p.db.ObservedContractsAddressList()
}

// ObservedContractAddressByType provides address of an observed contract by its type, if available.
func (p *Proxy) ObservedContractAddressByType(t string) *common.Address {
	adr, err := p.db.ObservedContractAddressByType(t)
	if err != nil {
		log.Errorf("contract lookup failed for %s, %s", t, err.Error())
		return nil
	}

	log.Noticef("%s contract is %s", t, adr.String())
	return adr
}

// NFTContractsTypeMap provides a map of observed contract addresses to corresponding
// contract type for ERC721 and ERC1155 contracts including their factory.
// In case of a factory contract, we need the deployed NFT type for processing.
func (p *Proxy) NFTContractsTypeMap() map[common.Address]string {
	return p.db.NFTContractsTypeMap()
}

// MinObservedBlockNumber provides the lowest observed block number.
func (p *Proxy) MinObservedBlockNumber(def uint64) uint64 {
	return p.db.MinObservedBlockNumber(def)
}

// loadObservedCollections loads observed collections into the repository.
func (p *Proxy) loadObservedCollections() {
	cl, err := p.shared.ObservedCollections()
	if err != nil {
		log.Critical("no observed collections; %s", err.Error())
		return
	}

	for _, adr := range cl {
		if !p.IsObservedContract(&adr) {
			p.addObservedCollection(&adr)
		}
	}
}

// addObservedCollection adds new observed NFT collection.
func (p *Proxy) addObservedCollection(adr *common.Address) {
	ct, err := p.NFTContractType(adr)
	if err != nil {
		log.Warningf("contract can not be observed; %s", err.Error())
		return
	}

	var blk uint64
	switch ct {
	case types.ContractTypeERC721:
		blk, err = p.Erc721StartingBlockNumber(adr)
	case types.ContractTypeERC1155:
		blk, err = p.Erc1155StartingBlockNumber(adr)
	default:
		err = fmt.Errorf("unknown contract type")
	}
	if err != nil {
		log.Warningf("contract %s can not be added; %s", adr.String(), err.Error())
		return
	}

	head, err := p.GetHeader(blk)
	if err != nil {
		log.Warningf("contract %s can not be added; %s", adr.String(), err.Error())
		return
	}

	// store observed contract into the repository
	p.AddObservedContract(&types.ObservedContract{
		Address:     *adr,
		Type:        ct,
		Created:     types.Time(time.Unix(int64(head.Time), 0)),
		Creator:     common.Address{},
		BlockNumber: blk,
		DeployedBy:  common.Hash{},
	})
}

// ContractTypeByAddress provides type of contract based on known address.
// We use pre-loaded NFT types map to perform this.
func (p *Proxy) ContractTypeByAddress(adr *common.Address) (string, error) {
	tp, ok := p.nftTypes[*adr]
	if !ok {
		return "", fmt.Errorf("address %s unknown", adr.String())
	}
	return tp, nil
}

// BasicContracts provides addresses of basic Artion contracts.
func (p *Proxy) BasicContracts() (c *types.Contracts) {
	return p.rpc.BasicContracts()
}

// IsEscrowContract test if the address is address of auction or other escrow contract.
func (p *Proxy) IsEscrowContract(addr common.Address) bool {
	return p.rpc.IsEscrowContract(addr)
}
