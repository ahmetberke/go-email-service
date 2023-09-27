package main

import (
	"fmt"
	"github.com/ahmetberke/go-email-service/configs"
	"github.com/ahmetberke/go-email-service/email"
	"github.com/ahmetberke/go-email-service/rabbit"
)

func main() {

	configurationManager := configs.NewConfigManager()
	rabbitConfig := configurationManager.GetRabbitConfig()
	queuesConfig := configurationManager.GetQueuesConfig()

	emailConsumer := email.Consumer{}

	rabbitClient := rabbit.NewRabbitClient(rabbitConfig, queuesConfig, emailConsumer)
	defer rabbitClient.CloseConnection()

	rabbitClient.DeclareExchangeQueueBindings()

	consumerChan := make(chan bool)
	rabbitClient.InitializeConsumers()
	fmt.Println("Started consumers")

	<-consumerChan
}
