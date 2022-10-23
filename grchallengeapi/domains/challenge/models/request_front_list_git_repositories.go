package challengemodels

type FrontListGitRepositoriesRequest struct {
	GitRepositoryFilterData
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}
