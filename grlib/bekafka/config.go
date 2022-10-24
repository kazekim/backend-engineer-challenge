package bekafka

type Config struct {
	BrokerUrl []string `json:"broker_url" mapstructure:"broker_url"`
}
