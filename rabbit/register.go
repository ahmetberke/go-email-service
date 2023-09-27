package rabbit

import (
	"errors"
	"github.com/ahmetberke/go-email-service/configs"
	"github.com/streadway/amqp"
)

type QueueConsumerMap map[configs.QueueConfig]func(delivery amqp.Delivery) error

var qcm QueueConsumerMap

func (c *Client) getRegisteredQueueConsumer() QueueConsumerMap {
	if qcm != nil {
		return qcm
	}
	queueConsumerMap := make(QueueConsumerMap)

	queueConsumerMap[c.queuesConfig.Email.EmailCreated] = c.emailConsumer.ConsumeEmailCreated

	return queueConsumerMap
}

func FindConsumer(routingKey string) (func(delivery amqp.Delivery) error, error) {
	for key, value := range qcm {
		if key.RoutingKey == routingKey {
			return value, nil
		}
	}
	return nil, errors.New("Consumer not found, Routing Key: " + routingKey)
}
