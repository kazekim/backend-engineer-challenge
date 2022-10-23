package grscankafka

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/bekafka"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

type Client interface {
	PublishStartGitRepositoryScanningMessage(serverId string, data PublishStartGitRepositoryScanningMessageParams) grerrors.Error
	StartSendCampaignWorker(action func(topic string, params PublishStartGitRepositoryScanningMessageParams) bekafka.WorkerStatus) grerrors.Error
}

type defaultClient struct {
	cfg *Config
	kc  bekafka.Client
}

//NewClient init new git repository scan kafka client
func NewClient(cfg *Config) Client {

	kCfg := bekafka.Config{
		BrokerUrl: []string{
			cfg.BrokerUrl,
		},
	}
	kc := bekafka.NewClient(&kCfg)

	return &defaultClient{
		cfg: cfg,
		kc:  kc,
	}
}
