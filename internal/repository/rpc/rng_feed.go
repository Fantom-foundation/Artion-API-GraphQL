// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// IsPendingRandomNumberRequest checks if the given random number request is pending.
func (o *Opera) IsPendingRandomNumberRequest(reqID *common.Hash, seed *common.Hash) bool {
	// do we have a connection to the RNG contract?
	if nil == o.rngFeedContract {
		log.Errorf("rng contract is not loaded")
		return false
	}

	req, err := o.rngFeedContract.GetRequest(nil, *reqID)
	if err != nil {
		log.Errorf("can not get the request %s; %s", reqID.String(), err.Error())
		return false
	}

	return 0 == bytes.Compare(req.Seed[:], (*seed)[:])
}

// FulfillRandomNumberRequest send the given random INT number into the RNG feed oracle
// as a response to the detected request.
func (o *Opera) FulfillRandomNumberRequest(reqID *common.Hash, rnd *big.Int) error {
	// do we have a connection to the RNG contract?
	if nil == o.rngFeedContract {
		return fmt.Errorf("rng contract is not loaded")
	}

	trx, err := o.rngFeedContract.FulfillRandomNumber(o.rngTransactor(), *reqID, rnd)
	if err != nil {
		log.Criticalf("could not feed rng to %s; %s", reqID.String(), err.Error())
		return err
	}

	// log pending transaction
	log.Infof("rng request %s resolved by %s", reqID.String(), trx.Hash().String())
	return nil
}

// rngTransactor provides signing transactor for the RNG request handler.
func (o *Opera) rngTransactor() *bind.TransactOpts {
	chain, err := hexutil.DecodeBig(cfg.RngOracle.ChainID)
	if err != nil {
		log.Criticalf("can not decode chain ID; %s", err.Error())
		return nil
	}

	tx, err := bind.NewKeyedTransactorWithChainID(&cfg.RngOracle.PrivateKey, chain)
	if err != nil {
		log.Criticalf("can not create signing authority; %s", err.Error())
		return nil
	}
	return tx
}
