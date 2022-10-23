package grwrouters

import (
	"fmt"
	"github.com/kazekim/backend-engineer-challenge/grlib/bekafka"
	"github.com/kazekim/backend-engineer-challenge/grlib/grscankafka"
)

func (r *defaultRouter) registerGitScannerWorker() {
	go func() {
		fmt.Println("start git scanner worker")
		_ = r.options.GitScannerMQClient.StartSendCampaignWorker(func(topic string, params grscankafka.PublishStartGitRepositoryScanningMessageParams) bekafka.WorkerStatus {

			vErr := r.cu.DoGitRepositoryScanningForResultId(params.ResultId)
			if vErr != nil {
				return bekafka.WorkerStatusFail
			}
			return bekafka.WorkerStatusSuccess
		})
	}()
}
