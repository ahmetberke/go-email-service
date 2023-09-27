package configs

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	configPath = "./resources"
	configType = "yaml"
)

type ConfigManager interface {
	GetRabbitConfig() RabbitConfig
	GetQueuesConfig() QueuesConfig
	GetSMTPConfig() SMTPConfig
}

type configManager struct {
	applicationConfig ApplicationConfig
	queuesConfig      QueuesConfig
}

func NewConfigManager() ConfigManager {
	env := os.Getenv("PROFILE")
	if env == "" {
		env = "local"
	}
	viper.AddConfigPath(configPath)
	viper.SetConfigType(configType)
	appConfig := readApplicationConf(env)
	queuesConfig := readQueuesConfig()
	return &configManager{applicationConfig: appConfig, queuesConfig: queuesConfig}
}

func (c configManager) GetRabbitConfig() RabbitConfig {
	return c.applicationConfig.Rabbit
}

func (c configManager) GetQueuesConfig() QueuesConfig {
	return c.queuesConfig
}

func (c configManager) GetSMTPConfig() SMTPConfig {
	return c.applicationConfig.SMTP
}

func readApplicationConf(env string) ApplicationConfig {
	viper.SetConfigName("application")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Couldn't load application configs, cannot start. Error details : %s", err.Error())
	}

	var conf ApplicationConfig
	c := viper.Sub(env)
	err = c.Unmarshal(&conf)
	if err != nil {
		log.Panicf("Configuration cannot deserialize. Terminating. Error details: %s", err.Error())
	}

	err = c.Unmarshal(&conf)
	if err != nil {
		log.Panicf("Configuration cannot deserialize. Terminating. Error details: %s", err.Error())
	}

	return conf

}

func readQueuesConfig() QueuesConfig {
	viper.SetConfigName("rabbit-queue")
	readConfigErr := viper.ReadInConfig()
	if readConfigErr != nil {
		log.Panicf("Couldn't load queues configuration, cannot start. Error details: %s", readConfigErr.Error())
	}
	var conf QueuesConfig
	c := viper.Sub("queue")
	unMarshalErr := c.Unmarshal(&conf)
	unMarshalSubErr := c.Unmarshal(&conf)
	if unMarshalErr != nil {
		log.Panicf("Configuration cannot deserialize. Terminating. Error details: %s", unMarshalErr.Error())
	}
	if unMarshalSubErr != nil {
		log.Panicf("Configuration cannot deserialize. Terminating. Error details: %s", unMarshalSubErr.Error())
	}
	return conf
}
