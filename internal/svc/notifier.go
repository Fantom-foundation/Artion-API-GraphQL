// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/types"
	"time"
)

const (
	// maxNotificationAgeDelivery represent the max age of a notification to be pushed to delivery service.
	maxNotificationAgeDelivery = -48 * time.Hour

	// notificationQueueTerminateDelay is the delay we let the queue clear out before forced termination.
	notificationQueueTerminateDelay = 2 * time.Second
)

// represents a function callback for notification delivery
type notificationDeliveryCallback func(*types.Notification, *types.NotificationTemplate) error

// notificationProcessor represents a service responsible for processing notifications
// of some events on NFT collections, tokens and social media links.
type notificationProcessor struct {
	// mgr represents the Manager instance
	mgr *Manager

	// sigStop represents the signal for closing the observer
	sigStop chan bool

	// inNotify represents the channel used by the service to receive new notification requests
	inNotify chan types.Notification

	// providers represents a map of notification delivery providers to their respective callback
	// delivery functions.
	providers map[string]notificationDeliveryCallback
}

// newNotificationProcessor creates a new instance of the event notification processor service.
func newNotificationProcessor(mgr *Manager) *notificationProcessor {
	return &notificationProcessor{
		mgr:     mgr,
		sigStop: make(chan bool, 1),
		providers: map[string]notificationDeliveryCallback{
			"sendgrid": repo.SendEmailNotificationBySendGrid,
		},
	}
}

// name provides the name fo the log observer service.
func (nop *notificationProcessor) name() string {
	return "notifier"
}

// init configures the notifier and subscribes it with the manager.
func (nop *notificationProcessor) init() {
	// link channels
	nop.inNotify = repo.NewNotifications()

	// add and run
	nop.mgr.add(nop)
}

// close signals the notifier service to terminate.
func (nop *notificationProcessor) close() {
	// delay termination to let the queue clear
	go func() {
		<-time.After(notificationQueueTerminateDelay)
		nop.sigStop <- true
	}()
}

// run collects incoming notification requests and processes them using
// appropriate service providers.
func (nop *notificationProcessor) run() {
	defer func() {
		nop.mgr.closed(nop)
	}()

	for {
		select {
		case <-nop.sigStop:
			return
		case evt, ok := <-nop.inNotify:
			if !ok {
				log.Noticef("no more notification are coming")
				return
			}
			nop.process(&evt)
		}
	}
}

// process the given notification request, if it has not been processed before.
func (nop *notificationProcessor) process(no *types.Notification) {
	// store the notification in database
	isNew, err := repo.StoreNotification(no)
	if err != nil {
		log.Errorf("notification #%d on %s/%s to %s not stored; %s",
			no.Type,
			no.Contract.String(),
			no.TokenId.String(),
			no.Recipient.String(),
			err.Error())
	}
	if !isNew {
		return
	}

	// check age; the Duration is negative, we are actually subtracting it from Now()
	if time.Now().UTC().Add(maxNotificationAgeDelivery).After(time.Time(no.TimeStamp)) {
		log.Warningf("notification #%d on %s/%s to %s is too old for delivery",
			no.Type,
			no.Contract.String(),
			no.TokenId.String(),
			no.Recipient.String())
		return
	}

	// pull applicable notification templates
	tmp := repo.NotificationTemplates(no.Type, &no.Contract, &no.TokenId)
	if tmp == nil || 0 == len(tmp) {
		log.Warningf("no templates for notification #%d on %s/%s to %s",
			no.Type,
			no.Contract.String(),
			no.TokenId.String(),
			no.Recipient.String())
		return
	}

	// loop all templates and deliver them
	for _, t := range tmp {
		provider, ok := nop.providers[t.Provider]
		if !ok {
			log.Warningf("no %s provider to deliver #%d on %s/%s to %s",
				t.Provider,
				no.Type,
				no.Contract.String(),
				no.TokenId.String(),
				no.Recipient.String())
			continue
		}

		err := provider(no, &t)
		if err != nil {
			log.Errorf("failed to deliver #%d on %s/%s to %s via %s; %s",
				no.Type,
				no.Contract.String(),
				no.TokenId.String(),
				no.Recipient.String(),
				t.Provider,
				err.Error(),
			)
		}
	}
}
