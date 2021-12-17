// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/types"
	"bytes"
	"github.com/ethereum/go-ethereum/common"
)

// notificationQueueCapacity represents the maximal capacity of notificationQueue queued to be sent.
const notificationQueueCapacity = 100

// pierreGaslyNFTCollectionAddress is the hardcoded address of the Pierre Gasly NFT collection.
var pierreGaslyNFTCollectionAddress = common.HexToAddress("0x475631Dbd805F46BE62D8F87a4f07CA8aFaF7E45")

// NewNotifications provides a channel for new notification requests queue.
func (p *Proxy) NewNotifications() chan types.Notification {
	return p.notificationQueue
}

// QueueNotificationForProcessing puts the given notification into the queue for async processing.
func (p *Proxy) QueueNotificationForProcessing(no *types.Notification) {
	// validate the notification before pushing it
	if p.NotificationCanSend(no) {
		p.notificationQueue <- *no
	}
}

// NotificationCanSend checks if the given notification can be sent.
func (p *Proxy) NotificationCanSend(no *types.Notification) bool {
	// Pierre Gasly burns are always passed; we need this one for now
	// @todo Remove hardcoded Pierre Gasly exception from notification check.
	if no.Type == types.NotifyNFTBurned && no.Contract != nil && 0 == bytes.Compare(no.Contract.Bytes(), pierreGaslyNFTCollectionAddress.Bytes()) {
		return true
	}

	// get the recipient config
	nst, err := p.GetNotificationSettings(no.Recipient)
	if err != nil {
		log.Errorf("can get notification settings for %s; %s", no.Recipient.String(), err.Error())
		return false
	}
	if nst == nil {
		log.Debugf("recipient %s does not have set notifications settings", no.Recipient.String())
		return false
	}

	rs, err := nst.IsTypeEnabled(no.Type)
	if err != nil {
		log.Errorf("notification setting not available for %s; %s", no.Recipient.String(), err.Error())
		return false
	}
	return rs
}
