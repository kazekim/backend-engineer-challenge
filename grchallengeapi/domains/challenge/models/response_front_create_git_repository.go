package challengemodels

import "github.com/kazekim/backend-engineer-challenge/grlib/grmodels"

type FrontCreateGitRepositoryResponse struct {
	Data grmodels.GitRepository `json:"data"`
}
