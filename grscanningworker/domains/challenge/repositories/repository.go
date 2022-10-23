package challengerepositories

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
	challengemodels "github.com/kazekim/backend-engineer-challenge/grscanningworker/domains/challenge/models"
)

type Repository interface {
	GetGitRepositoryScanResultById(id string) (*grmodels.GitRepositoryScanResultWithDetail, grerrors.Error)
	UpdateGitRepositoryScanResultById(id string, data challengemodels.UpdateGitRepositoryScanResultData) grerrors.Error
}
