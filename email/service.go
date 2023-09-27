package email

import (
	"crypto/tls"
	"github.com/ahmetberke/go-email-service/configs"
	"gopkg.in/gomail.v2"
	"log"
)

type EmailService struct {
	config configs.SMTPConfig
}

func NewService(config configs.SMTPConfig) EmailService {
	return EmailService{config: config}
}

func (es *EmailService) SendEmail(email Email) error {

	m := gomail.NewMessage()

	m.SetHeader("From", es.config.User)
	m.SetHeader("To", email.Recipient)
	m.SetHeader("Subject", email.Subject)

	m.SetBody("text/plain", email.Content)

	d := gomail.NewDialer(es.config.Host, es.config.Port, es.config.User, es.config.Password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	log.Printf("--> Email Sent <--")
	return nil

}
