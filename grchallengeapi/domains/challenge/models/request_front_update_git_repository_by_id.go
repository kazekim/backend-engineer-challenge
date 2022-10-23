package challengemodels

type FrontUpdateGitRepositoryByIdRequest struct {
	RepositoryId string `path:"repository_id"`
	UpdateGitRepositoryData
}
