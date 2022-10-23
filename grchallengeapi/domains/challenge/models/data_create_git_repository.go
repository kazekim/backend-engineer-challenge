package challengemodels

type CreateGitRepositoryData struct {
	Name string `json:"name" binding:"required"`
	Url  string `json:"url" binding:"required,url"`
}
