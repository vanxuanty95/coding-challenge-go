package helper

import (
	"coding-challenge-go/cmd/api/config"
	"coding-challenge-go/pkg/api/dictionary"
	"coding-challenge-go/pkg/logger"
)

const (
	SMS_TYPE   = "sms"
	EMAIL_TYPE = "email"
)

type NotifiersFactory struct {
	gdgLogger        logger.Logger
	listAllNotifiers []Notifier
}

func NewNotifiersFactory(cfg *config.Config) *NotifiersFactory {
	listAllNotifiers := CreateNotifiers(cfg)
	return &NotifiersFactory{
		gdgLogger:        logger.WithPrefix("notifiers-factory"),
		listAllNotifiers: listAllNotifiers,
	}
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

func (nf *NotifiersFactory) SendNotification(info NotificationsInfo) {
	for _, currNotifier := range nf.listAllNotifiers {
		err := currNotifier.StockChanged(info)
		if err != nil {
			nf.gdgLogger.Errorln(dictionary.SendNotificationError, err)
		}
	}
}
