package challengemodels

import grcmodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/models"

type FrontListGitRepositoriesResponse struct {
	Datas []grcmodels.GitRepository `json:"datas"`
	Page int64 `json:"page"`
	Count int64 `json:"count"`
}
