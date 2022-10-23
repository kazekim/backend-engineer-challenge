package challengemodels

import grcmodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/models"

type FrontStartGitRepositoryScanningResponse struct {
	Data grcmodels.GitRepositoryScanResult `json:"data"`
}
