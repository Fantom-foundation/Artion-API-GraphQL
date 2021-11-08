// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/repository/email"
	"artion-api-graphql/internal/types"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"strings"
)

// notificationQueueCapacity represents the maximal capacity of notificationQueue queued to be sent.
const notificationQueueCapacity = 100

// NewNotifications provides a channel for new notification requests queue.
func (p *Proxy) NewNotifications() chan types.Notification {
	return p.notificationQueue
}

// QueueNotificationForProcessing puts the given notification into the queue for async processing.
func (p *Proxy) QueueNotificationForProcessing(no *types.Notification) {
	p.notificationQueue <- *no
}

// StoreNotification stores the given notification in persistent storage.
// The function returns true if the notification was unique and didn't exist before.
func (p *Proxy) StoreNotification(no *types.Notification) (bool, error) {
	return p.db.StoreNotification(no)
}

// NotificationTemplates pulls a list of notification templates applicable to the
func (p *Proxy) NotificationTemplates(nt int32, contract *common.Address, tokenID *hexutil.Big) []types.NotificationTemplate {
	return p.db.NotificationTemplates(nt, contract, tokenID)
}

// SendEmailNotificationBySendGrid represents an email notification sender using SendGrid API.
func (p *Proxy) SendEmailNotificationBySendGrid(no *types.Notification, nt *types.NotificationTemplate) error {
	if no == nil || nt == nil {
		return fmt.Errorf("missing notification and/or template")
	}

	// collect recipient user involved
	user, err := p.GetUser(no.Recipient)
	if err != nil {
		return err
	}
	if user == nil {
		return fmt.Errorf("user not found at %s", no.Recipient.String())
	}

	// assign recipient
	recipient := nt.Recipient
	if recipient == nil && nil != user.Email && strings.TrimSpace(*user.Email) != "" {
		recipient = user.Email
	}

	// any email set on User
	if nil == recipient {
		log.Warningf("no recipient on user %s", no.Recipient.String())
		return nil
	}

	// collect address
	addr, err := p.GetShippingAddress(no.Recipient)
	if err != nil {
		return err
	}

	// split recipient to CC fields
	fields := strings.Fields(*recipient)
	cc := make([]*mail.Email, len(fields))
	for i, eml := range fields {
		cc[i] = mail.NewEmail("", eml)
	}

	// send the email
	err = email.SendGridDeliverDynamicTemplate(
		mail.NewEmail(nt.SenderName, nt.SenderID),
		cc[0],
		cc[1:],
		nt.TemplateID,
		nt.Subject,
		dynamicTemplateData(no, user, addr, nt.ExtendedParams),
	)
	if err != nil {
		log.Errorf("email notification failed; %s", err.Error())
		return err
	}
	return nil
}

// dynamicTemplateData creates a map of key->value dynamic data points from provided
// source elements.
func dynamicTemplateData(no *types.Notification, usr *types.User, ship *types.ShippingAddress, ext *string) map[string]interface{} {
	var list map[string]interface{}

	// do we have an ext?
	if nil != ext {
		err := json.Unmarshal([]byte(*ext), &list)
		if err != nil {
			log.Errorf("invalid extended data on template; %s", err.Error())
		}
	}

	// add notification details
	if no != nil {
		list["collection"] = no.Contract.String()
		list["contract"] = no.Contract.String()
		list["token_id"] = no.TokenId.String()

		log.Infof("NFT collection %s, token ID %s", no.Contract.String(), no.TokenId.String())
	}

	// add user
	if usr != nil {
		list["account"] = usr.Address.String()
		list["address"] = usr.Address.String()
		list["email"] = *usr.Email
		list["alias"] = ""

		if nil != usr.Username {
			list["alias"] = *usr.Username
		}
		log.Infof("user %s, email address %s", list["alias"], list["email"])
	}

	// add shipping address
	if ship != nil {
		list["shipping_name"] = ship.FullName
		list["shipping_phone"] = ship.Phone
		list["shipping_street"] = ship.Street
		list["shipping_apartment"] = ship.Apartment
		list["shipping_city"] = ship.City
		list["shipping_state"] = ship.State
		list["shipping_country"] = ship.Country
		list["shipping_zip"] = ship.Zip

		log.Infof("shipping address set as %s, %s", ship.City, ship.Country)
	}
	return list
}
