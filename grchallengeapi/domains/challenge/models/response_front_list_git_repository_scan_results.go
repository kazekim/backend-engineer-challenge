package challengemodels

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
)

type FrontListGitRepositoryScanResultsResponse struct {
	Datas []grmodels.GitRepositoryScanResultWithDetail `json:"datas"`
	Page  int64                                        `json:"page"`
	Count int64                                        `json:"count"`
}
