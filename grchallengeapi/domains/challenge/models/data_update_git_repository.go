package challengemodels

type UpdateGitRepositoryData struct {
	Name     *string `form:"name"`
	Url      *string `form:"url" binding:"url"`
	IsActive *bool   `form:"is_active"`
}
