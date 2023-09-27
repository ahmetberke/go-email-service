package email

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

type Consumer struct {
	emailService EmailService
}

func NewConsumer(service EmailService) Consumer {
	return Consumer{
		emailService: service,
	}
}

func (c *Consumer) ConsumeEmailCreated(delivery amqp.Delivery) error {

	var email Email
	err := json.Unmarshal(delivery.Body, &email)
	if err != nil {
		return err
	}

	err = c.emailService.SendEmail(email)
	if err != nil {
		return err
	}
	fmt.Printf("Email; Recipient : %s, Content : %s", email.Recipient, email.Content)

	return nil
}
