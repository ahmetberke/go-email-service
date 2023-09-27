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
	config := readApplicationConf(env)
	return &configManager{applicationConfig: config}
}

func (c configManager) GetRabbitConfig() RabbitConfig {
	return c.applicationConfig.Rabbit
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
