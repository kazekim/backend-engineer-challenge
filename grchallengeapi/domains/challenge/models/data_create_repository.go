package challengemodels

type CreateRepositoryData struct {
	Name string `json:"name" binding:"required"`
	Url string `json:"url" binding:"required,url"`
}
