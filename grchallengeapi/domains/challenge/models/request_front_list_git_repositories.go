package challengemodels

type FrontListGitRepositoriesRequest struct {
	GitRepositoryFilterData
	Page  int64 `form:"page"`
	Limit int64 `form:"limit"`
}
