// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"crypto/rand"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// requestedRandomNumber handles log event for Random Number Oracle request.
// RandomNumberOracle::RandomNumberRequested(bytes32 requestID, bytes32 seed)
func requestedRandomNumber(evt *eth.Log, _ *logObserver) {
	// sanity check: 1 + 0 topics; 2 x bytes32 = 2 x 32 bytes of data = 64 bytes
	if len(evt.Data) != 64 || len(evt.Topics) != 1 {
		log.Errorf("not RandomNumberOracle::RandomNumberRequested() event #%d/#%d; expected 64 bytes of data, %d given; expected 1 topic, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	// extract the request ID we need to do
	requestID := common.BytesToHash(evt.Data[:32])
	seed := common.BytesToHash(evt.Data[32:])
	if !repo.IsPendingRandomNumberRequest(&requestID, &seed) {
		log.Noticef("rng request %s already done", requestID.String())
		return
	}

	// limit the generated number to 128 bits - 1
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(128), nil).Sub(max, big.NewInt(1))

	// generate random number between 0 and max
	rnd, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Errorf("could not create random number; %s", err.Error())
		return
	}

	// try to fulfill
	if err := repo.FulfillRandomNumberRequest(&requestID, rnd); err != nil {
		log.Errorf("could not feed random number for %s; %s", requestID.String(), err.Error())
		return
	}

	log.Infof("rng request %s fulfilled with %s", requestID.String(), (*hexutil.Big)(rnd).String())
}
