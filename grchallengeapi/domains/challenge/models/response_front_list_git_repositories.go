package challengemodels

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
)

type FrontListGitRepositoriesResponse struct {
	Datas []grmodels.GitRepository `json:"datas"`
	Page  int64                    `json:"page"`
	Count int64                    `json:"count"`
}
