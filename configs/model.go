package configs

type ApplicationConfig struct {
	Rabbit RabbitConfig `yaml:"rabbit"`
	SMTP   SMTPConfig   `yaml:"smtp"`
}

type RabbitConfig struct {
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	VirtualHost    string `yaml:"virtualHost"`
	ConnectionName string `yaml:"connectionName"`
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
}

type SMTPConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type QueuesConfig struct {
	Email EmailQueueConfig `yaml:"email"`
}

type EmailQueueConfig struct {
	EmailCreated QueueConfig `yaml:"emailCreated"`
}

type QueueConfig struct {
	PrefetchCount int    `yaml:"prefetchCount"`
	ChannelCount  int    `yaml:"channelCount"`
	Exchange      string `yaml:"exchange"`
	ExchangeType  string `yaml:"exchangeType"`
	RoutingKey    string `yaml:"routingKey"`
	Queue         string `yaml:"queue"`
}
