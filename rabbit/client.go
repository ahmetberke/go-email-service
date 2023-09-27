package rabbit

import (
	"fmt"
	"github.com/ahmetberke/go-email-service/configs"
	"github.com/ahmetberke/go-email-service/email"
	"github.com/streadway/amqp"
	"log"
	"time"
)

type Client struct {
	connection    *amqp.Connection
	queuesConfig  configs.QueuesConfig
	emailConsumer email.Consumer
}

func NewRabbitClient(rabbitConfig configs.RabbitConfig, queuesConfig configs.QueuesConfig, emailConsumer email.Consumer) *Client {
	return &Client{
		connection:    createConnection(rabbitConfig),
		queuesConfig:  queuesConfig,
		emailConsumer: emailConsumer,
	}
}

func createConnection(rabbitConfig configs.RabbitConfig) *amqp.Connection {
	amqpConfig := amqp.Config{
		Properties: amqp.Table{
			"connection_name": rabbitConfig.ConnectionName,
		},
		Heartbeat: 30 * time.Second,
	}
	connectionURL := getConnectionURL(rabbitConfig)
	connection, err := amqp.DialConfig(connectionURL, amqpConfig)
	if err != nil {
		_ = connection.Close()
		log.Panicf("Client cannot deserialize. Terminating. Error details: %s", err.Error())
	}
	log.Printf("RabbitMQ connected. Host: %s, Port: %s, Virtual Host: %s", rabbitConfig.Host, rabbitConfig.Port, rabbitConfig.VirtualHost)
	return connection
}

func getConnectionURL(config configs.RabbitConfig) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/%s", config.Username, config.Password, config.Host, config.Port, config.VirtualHost)
}

func (c *Client) DeclareExchangeQueueBindings() {
	channel := c.CreateChannel(0)
	configs := c.getRegisteredQueueConsumer()
	for queueConfig, _ := range configs {
		declareExchange(channel, queueConfig)
		declareQueue(channel, queueConfig)
		declareDeadLetterQueue(channel, queueConfig)
		bindQueue(channel, queueConfig)
		err := channel.Qos(queueConfig.PrefetchCount, 0, false)
		if err != nil {
			log.Panicf("PrefetchCount could not defined. Terminating. Error details: %s", err.Error())
		}
	}
}

func (c *Client) CreateChannel(prefetchCount int) *amqp.Channel {
	channel, err := c.connection.Channel()
	if err != nil {
		_ = channel.Close()
		log.Panicf("Channel couldn't created. Terminating. Error details: %s", err.Error())
	}
	err = channel.Qos(prefetchCount, 0, false)
	if err != nil {
		log.Panicf("PrefetchCount couldn't defined. Terminating. Error details: %s", err.Error())
	}
	return channel
}

func declareExchange(channel *amqp.Channel, queueConfig configs.QueueConfig) {
	err := channel.ExchangeDeclare(queueConfig.Exchange, queueConfig.ExchangeType, true, false, false, false, nil)
	if err != nil {
		log.Panicf("Exchange could not declared. Terminating. Error details: %s", err.Error())
	}
}
func declareQueue(channel *amqp.Channel, queueConfig configs.QueueConfig) {
	deadLetterArgs := getDeadLetterArgs(queueConfig.Queue)
	_, err := channel.QueueDeclare(queueConfig.Queue, true, false, false, false, deadLetterArgs)
	if err != nil {
		log.Panicf("Queue could not declared. Terminating. Error details: %s", err.Error())
	}
}

func declareDeadLetterQueue(channel *amqp.Channel, queueConfig configs.QueueConfig) {
	_, err := channel.QueueDeclare(queueConfig.Queue+".deadLetter", true, false, false, false, nil)
	if err != nil {
		log.Panicf("Queue could not declared. Terminating. Error details: %s", err.Error())
	}
}

func bindQueue(channel *amqp.Channel, queueConfig configs.QueueConfig) {
	err := channel.QueueBind(queueConfig.Queue, queueConfig.RoutingKey, queueConfig.Exchange, false, nil)
	if err != nil {
		log.Panicf("Binding could not defined. Terminating. Error details: %s", err.Error())
	}
}

func getDeadLetterArgs(queueName string) amqp.Table {
	return amqp.Table{
		"x-dead-letter-exchange":    "",
		"x-dead-letter-routing-key": queueName + ".deadLetter",
	}
}

func (c *Client) CloseConnection() {
	c.connection.Close()
}
