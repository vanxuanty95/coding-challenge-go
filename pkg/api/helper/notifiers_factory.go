package helper

import (
	"coding-challenge-go/cmd/api/config"
	"coding-challenge-go/pkg/api/dictionary"
	"coding-challenge-go/pkg/logger"
	"context"
	"fmt"
	"time"
)

const (
	SMS_TYPE   = "sms"
	EMAIL_TYPE = "email"

	MAX_WORKER = 10
)

type NotifiersFactory struct {
	gdgLogger        logger.Logger
	listAllNotifiers []Notifier
	listMessage      chan NotificationsInfo
	stop             chan bool
}

func NewNotifiersFactory(cfg *config.Config) *NotifiersFactory {
	listAllNotifiers := CreateNotifiers(cfg)
	noiFac := &NotifiersFactory{
		gdgLogger:        logger.WithPrefix("notifiers-factory"),
		listAllNotifiers: listAllNotifiers,
		listMessage:      make(chan NotificationsInfo, MAX_WORKER),
		stop:             make(chan bool, 1),
	}
	noiFac.start()
	return noiFac
}

func CreateNotifiers(cfg *config.Config) (listAllNotifiers []Notifier) {
	for _, currType := range cfg.Seller.Notification.Type {
		switch currType {
		case SMS_TYPE:
			listAllNotifiers = append(listAllNotifiers, NewSmsProvider(cfg))
		case EMAIL_TYPE:
			listAllNotifiers = append(listAllNotifiers, NewEmailProvider(cfg))
		}
	}
	return listAllNotifiers
}

func (nf *NotifiersFactory) start() {
	go func() {
		for {
			select {
			case message := <-nf.listMessage:
				nf.SendNotification(message)
			case <-nf.stop:
				return
			}
		}
	}()
}

func (nf *NotifiersFactory) SendChannel(info NotificationsInfo) {
	for {
		if len(nf.listMessage) > MAX_WORKER {
			nf.gdgLogger.Errorln("not worker free")
		} else {
			nf.listMessage <- info
			break
		}
	}
}

func (nf *NotifiersFactory) SendNotification(info NotificationsInfo) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	okay := make(chan bool, 1)

	go func(okay *chan bool) {
		time.Sleep(3 * time.Second)
		for _, currNotifier := range nf.listAllNotifiers {
			{
				err := currNotifier.StockChanged(info)
				if err != nil {
					nf.gdgLogger.Errorln(dictionary.SendNotificationError, err)
				}
				*okay <- true
			}
		}
	}(&okay)

	select {
	case <-ctx.Done():
		fmt.Println("timeout")
	case <-okay:
		break
	}
}
