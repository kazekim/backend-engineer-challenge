package challengemodels

type FrontListGitRepositoryScanResultsRequest struct {
	GitRepositoryScanResultFilterData
	Page  int64 `form:"page"`
	Limit int64 `form:"limit"`
}
