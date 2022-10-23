package bekafka

import "github.com/kazekim/backend-engineer-challenge/grlib/grerrors"

type Client interface {
	NewProducer() (*Producer, grerrors.Error)
	NewWorker() (*Worker, grerrors.Error)
}

type defaultClient struct {
	cfg *Config
}

//NewClient init bekafka client
func NewClient(cfg *Config) Client {
	return &defaultClient{
		cfg: cfg,
	}
}
