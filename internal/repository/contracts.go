// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

// AddObservedContract adds new observed contract into the persistent storage.
func (p *Proxy) AddObservedContract(oc *types.ObservedContract) {
	err := p.db.AddObservedContract(oc)
	if err != nil {
		log.Criticalf("can not add observed contract %s; %s", oc.Address.String(), err.Error())
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
	blk := p.db.MinObservedBlockNumber(def)

	col, err := p.ObservedCollections()
	if err != nil {
		log.Warningf("could not account for collections; %s", err.Error())
		return blk
	}

	var bn uint64
	for _, adr := range col {
		ct, err := p.NFTContractType(&adr)
		if err != nil {
			continue
		}

		switch ct {
		case types.ContractTypeERC721:
			bn, err = p.Erc721StartingBlockNumber(&adr)
		case types.ContractTypeERC1155:
			bn, err = p.Erc1155StartingBlockNumber(&adr)
		}
		if err != nil {
			continue
		}

		if bn < blk {
			blk = bn
		}
	}

	return blk
}

// ObservedCollections provides list of addresses of observed collections known to Artion API.
func (p *Proxy) ObservedCollections() ([]common.Address, error) {
	return p.shared.ObservedCollections()
}
