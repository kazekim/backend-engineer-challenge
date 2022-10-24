package grenums

type ScanStatus string

const (
	ScanStatusQueued     ScanStatus = "Queued"
	ScanStatusInProgress ScanStatus = "In Progress"
	ScanStatusSuccess    ScanStatus = "Success"
	ScanStatusFailure    ScanStatus = "Failure"
)
