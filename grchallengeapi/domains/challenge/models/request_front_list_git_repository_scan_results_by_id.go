package challengemodels

type FrontListGitRepositoryScanResultsRequest struct {
	GitRepositoryScanResultFilterData
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}
