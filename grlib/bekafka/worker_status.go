package bekafka

type WorkerStatus string

const (
	WorkerStatusSuccess   WorkerStatus = "success"
	WorkerStatusTerminate WorkerStatus = "terminate"
	WorkerStatusFail      WorkerStatus = "fail"
)
