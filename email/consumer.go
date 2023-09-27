package email

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Consumer struct {
}

func (c *Consumer) ConsumeEmailCreated(delivery amqp.Delivery) error {
	payload := string(delivery.Body)
	fmt.Printf("ConsumeEmailCreated %s \n", payload)
	return nil
}
