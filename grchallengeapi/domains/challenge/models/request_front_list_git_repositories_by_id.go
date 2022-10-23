package challengemodels

type FrontListGitRepositoriesByIdRequest struct {
	ListGitRepositoriesFilterData
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
}
