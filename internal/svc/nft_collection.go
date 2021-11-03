// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"github.com/ethereum/go-ethereum/common"
)

// nftCollectionValidator represents a service responsible for processing new NFT collections
// by validating them, adding them into the service observers and making sure new collections
// are properly scanned and registered.
type nftCollectionValidator struct {
	// mgr represents the Manager instance
	mgr *Manager

	// sigStop represents the signal for closing the observer
	sigStop chan bool

	// inCollections represents an input channel being fed with addresses
	// of newly registered collection contracts.
	inCollections chan common.Address
}

// newNFTCollectionValidator creates a new instance of the NFT validator.
func newNFTCollectionValidator(mgr *Manager) *nftCollectionValidator {
	return &nftCollectionValidator{
		mgr:     mgr,
		sigStop: make(chan bool, 1),
	}
}

// name provides the name fo the log observer service.
func (cv *nftCollectionValidator) name() string {
	return "collection validator"
}

// init configures the NFT collection validator and subscribes it with the manager.
func (cv *nftCollectionValidator) init() {
	// link channels
	cv.inCollections = repo.NewCollectionQueue()

	// add and run
	cv.mgr.add(cv)
}

// close signals the NFT collection validator to terminate.
func (cv *nftCollectionValidator) close() {
	cv.sigStop <- true
}

// run collects incoming NFT collections and processes them through.
func (cv *nftCollectionValidator) run() {
	defer func() {
		cv.mgr.closed(cv)
	}()

	for {
		select {
		case <-cv.sigStop:
			return
		case adr, ok := <-cv.inCollections:
			if !ok {
				return
			}

			if cv.isValid(&adr) {
				// push the address as a new observed contract
			}
		}
	}
}

// isValid checks if the address is a valid ERC-721 or ERC-1155 contract.
func (cv *nftCollectionValidator) isValid(adr *common.Address) bool {
	log.Infof("validating contract at %s", adr.String())

	return true
}
