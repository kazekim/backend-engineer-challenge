package grscankafka

type Config struct {
	BrokerUrl string `json:"kafka_broker_url" mapstructure:"kafka_broker_url"`
	ServerId  string `json:"kafka_server_id" mapstructure:"kafka_server_id"`
}
