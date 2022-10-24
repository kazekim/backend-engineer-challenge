package grscankafka

import (
	"encoding/json"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

// PublishStartGitRepositoryScanningMessage send git repository scanning job into message queue for waiting to start job
func (c *defaultClient) PublishStartGitRepositoryScanningMessage(serverId string, data PublishStartGitRepositoryScanningMessageParams) grerrors.Error {

	messageBytes, err := json.Marshal(data)
	if err != nil {
		return grerrors.NewKafkaProducerError(err)
	}
	messages := string(messageBytes)

	producer, vErr := c.kc.NewProducer()
	if vErr != nil {
		return vErr
	}
	defer producer.Close()

	vErr = producer.SendMessage(serverId, messages)
	if vErr != nil {
		return vErr
	}

	return nil
}
