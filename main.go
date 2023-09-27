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

	emailService := email.NewService(configurationManager.GetSMTPConfig())
	emailConsumer := email.NewConsumer(emailService)

	rabbitClient := rabbit.NewRabbitClient(rabbitConfig, queuesConfig, emailConsumer)
	defer rabbitClient.CloseConnection()

	rabbitClient.DeclareExchangeQueueBindings()

	consumerChan := make(chan bool)
	rabbitClient.InitializeConsumers()
	fmt.Println("Started consumers")

	<-consumerChan
}
