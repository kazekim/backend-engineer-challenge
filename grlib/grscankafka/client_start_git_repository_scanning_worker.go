package grscankafka

import (
	"encoding/json"
	"github.com/kazekim/backend-engineer-challenge/grlib/bekafka"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

// StartSendCampaignWorker start worker to receive job from message queue
func (c *defaultClient) StartSendCampaignWorker(action func(topic string, params PublishStartGitRepositoryScanningMessageParams) bekafka.WorkerStatus) grerrors.Error {

	worker, vErr := c.kc.NewWorker()
	if vErr != nil {
		return vErr
	}

	vErr = worker.ConsumePartition(c.cfg.ServerId, func(topic, message string) bekafka.WorkerStatus {

		var params PublishStartGitRepositoryScanningMessageParams
		err := json.Unmarshal([]byte(message), &params)
		if err != nil {
			return bekafka.WorkerStatusFail
		}
		return action(topic, params)
	})
	if vErr != nil {
		return vErr
	}
	return nil
}
