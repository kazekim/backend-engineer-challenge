package challengemodels

type FrontListGitRepositoryScanResultsByRepositoryIdRequest struct {
	RepositoryId string `path:"repository_id"`
	GitRepositoryScanResultFilterData
	Page  int64 `form:"page"`
	Limit int64 `form:"limit"`
}
