package helper

import (
	"coding-challenge-go/cmd/api/config"
	"fmt"
	"net/smtp"
	"strings"
)

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
)

func NewEmailProvider(cfg *config.Config) *EmailProvider {
	return &EmailProvider{cfg: cfg}
}

type EmailProvider struct {
	cfg *config.Config
}

func (ep *EmailProvider) StockChanged(info NotificationsInfo) error {
	from := ep.cfg.Seller.Notification.Template.Email.Sender.From
	password := ep.cfg.Seller.Notification.Template.Email.Sender.Password

	to := []string{info.SellerEmail}

	subject := fmt.Sprintf("%v%v", ep.cfg.Seller.Notification.Template.Email.Subject, "\n")

	body := strings.NewReplacer("{seller_name}", info.SellerName,
		"{product_name}", info.ProductName).Replace(ep.cfg.Seller.Notification.Template.Email.Body)
	msg := []byte(fmt.Sprintf("%v%v\n%v", subject, MIME, body))

	auth := smtp.PlainAuth("", from, password, ep.cfg.Seller.Notification.Template.Email.Sender.Host)

	err := smtp.SendMail(ep.cfg.Seller.Notification.Template.Email.Sender.Add, auth, from, to, msg)
	if err != nil {
		return err
	}

	return nil
}
