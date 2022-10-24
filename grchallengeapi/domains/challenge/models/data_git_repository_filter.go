package challengemodels

type GitRepositoryFilterData struct {
	Id       *string `form:"id"`
	Name     *string `form:"name"`
	IsActive *bool   `form:"is_active"`
}
