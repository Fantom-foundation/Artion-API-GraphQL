// Package rpc provides high level access to the Fantom Opera blockchain
// node through RPC interface.
package rpc

import (
	"artion-api-graphql/internal/types"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

// ExtendAuctionDetailAt adds contract stored details to the provided auction record.
func (o *Opera) ExtendAuctionDetailAt(au *types.Auction, block *big.Int) error {
	// get auction details
	res, err := o.auctionContract.GetAuction(&bind.CallOpts{
		BlockNumber: block,
		Context:     context.Background(),
	}, au.Contract, (*big.Int)(&au.TokenId))
	if err != nil {
		log.Errorf("auction %s/%s not available; %s",
			au.Contract.String(), au.TokenId.String(), au.Owner.String(), err.Error())
		return err
	}

	// make sure we have what we came for
	if nil == res.ReservePrice || nil == res.StartTime || nil == res.EndTime {
		return fmt.Errorf("missing mandatory field on auction %s/%s", au.Contract.String(), au.TokenId.String())
	}

	// transfer values
	au.Owner = res.Owner
	au.ReservePrice = (hexutil.Big)(*res.ReservePrice)
	au.StartTime = types.Time(time.Unix(res.StartTime.Int64(), 0))
	au.EndTime = types.Time(time.Unix(res.EndTime.Int64(), 0))

	return nil
}
