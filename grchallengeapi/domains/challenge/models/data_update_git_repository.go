package challengemodels

type UpdateGitRepositoryData struct {
	Name string `json:"name" binding:"required"`
	Url  string `json:"url" binding:"required,url"`
}
