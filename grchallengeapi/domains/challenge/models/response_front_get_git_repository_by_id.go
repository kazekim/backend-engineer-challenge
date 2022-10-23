package challengemodels

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
)

type FrontGetGitRepositoryByIdResponse struct {
	Data grmodels.GitRepository `json:"data"`
}
