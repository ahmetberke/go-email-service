package configs

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	configPath = "./resources"
	configType = "yaml"
	configName = "application"
)

type ConfigManager interface {
	GetRabbitConfig() RabbitConfig
}

func NewConfigManager() ConfigManager {
	env := os.Getenv("PROFILE")
	if env == "" {
		env = "local"
	}
	viper.AddConfigPath(configPath)
	viper.SetConfigType(configType)
	viper.SetConfigName(configName)
	config := readConf(env)
	return &config
}

type Config struct {
	Rabbit RabbitConfig `yaml:"rabbit"`
}

type RabbitConfig struct {
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	VirtualHost    string `yaml:"virtualHost"`
	ConnectionName string `yaml:"connectionName"`
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
}

func (c Config) GetRabbitConfig() RabbitConfig {
	return c.Rabbit
}

func readConf(env string) Config {
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Couldn't load application configs, cannot start. Error details : %s", err.Error())
	}

	var conf Config
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
