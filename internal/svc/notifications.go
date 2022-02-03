// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// maxFollowerNotificationSize defines the max size of followers pool we handle with notifications
const maxFollowerNotificationSize = 50

// notifyEventToOwner queues token notification to the NFT owner.
func notifyEventToOwner(ntt int32, contract common.Address, tokenID hexutil.Big, owner common.Address, from *common.Address, ts types.Time) {
	repo.QueueNotificationForProcessing(&types.Notification{
		Type:       ntt,
		Contract:   &contract,
		TokenId:    &tokenID,
		TimeStamp:  ts,
		Recipient:  owner,
		Originator: from,
	})
}

// notifyEventToFollowers notifies followers of an NFT owner about an event.
func notifyEventToFollowers(ntt int32, contract common.Address, tokenID hexutil.Big, owner common.Address, ts types.Time) {
	list := repo.MustFollowers(owner)
	if len(list) > maxFollowerNotificationSize {
		log.Criticalf("too many followers on %s/%s, owner %s [%d], notifications #%d not sent", contract.String(), tokenID.String(), owner.String(), len(list), ntt)
		return
	}

	// queue notifications
	for _, fo := range list {
		repo.QueueNotificationForProcessing(&types.Notification{
			Type:       ntt,
			Contract:   &contract,
			TokenId:    &tokenID,
			TimeStamp:  ts,
			Recipient:  fo,
			Originator: &owner,
		})
	}
}
