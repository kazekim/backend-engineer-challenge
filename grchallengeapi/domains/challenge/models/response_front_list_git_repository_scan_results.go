package challengemodels

import grcmodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/models"

type FrontListGitRepositoryScanResultsResponse struct {
	Datas []grcmodels.GitRepositoryScanResult `json:"datas"`
	Page int64 `json:"page"`
	Count int64 `json:"count"`
}
