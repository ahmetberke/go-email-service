package email

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

type Consumer struct {
}

func (c *Consumer) ConsumeEmailCreated(delivery amqp.Delivery) error {

	var email Email
	err := json.Unmarshal(delivery.Body, &email)
	if err != nil {
		return err
	}

	fmt.Printf("Email; Recipient : %s, Content : %s", email.Recipient, email.Content)

	return nil
}
