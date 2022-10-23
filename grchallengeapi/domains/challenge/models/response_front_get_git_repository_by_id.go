package challengemodels

import grcmodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/models"

type FrontGetGitRepositoryByIdResponse struct {
	Data grcmodels.GitRepository `json:"data"`
}
