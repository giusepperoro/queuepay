package config

type ServiceConfiguration struct {
	PostgresConnectUrl string `yaml:"postgresConnectUrl"`
	RabbitMQConnectUrl string `yaml:"rabbitMQConnectUrl"`
	ServerAddressUrl   string `yaml:"serverAddressUrl"`
}
