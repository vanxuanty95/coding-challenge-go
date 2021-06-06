package helper

import (
	"coding-challenge-go/cmd/api/config"
	"coding-challenge-go/pkg/logger"
	"strings"
)

func NewSmsProvider(cfg *config.Config) *SmsProvider {
	return &SmsProvider{
		gdgLogger: logger.WithPrefix("sms-provider"),
		cfg:       cfg,
	}
}

type SmsProvider struct {
	gdgLogger logger.Logger
	cfg       *config.Config
}

func (sp *SmsProvider) StockChanged(info NotificationsInfo) error {
	body := strings.NewReplacer("{seller_UUID}", info.SellerUUID, "{seller_Phone}", info.SellerPhone,
		"{product_name}", info.ProductName).Replace(sp.cfg.Seller.Notification.Template.Sms)
	sp.gdgLogger.Infoln(body)
	return nil
}
