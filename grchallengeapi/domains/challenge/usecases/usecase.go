package challengeusecases

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
)

type UseCase interface {
	CreateGitRepository(data challengemodels.CreateGitRepositoryData) (*grmodels.GitRepository, grerrors.Error)
	DeleteGitRepositoryById(id string) grerrors.Error
	GetGitRepositoryById(id string) (*grmodels.GitRepository, grerrors.Error)
	ListGitRepositories(filter challengemodels.GitRepositoryFilterData, values ...int64) (*[]grmodels.GitRepository, int64, grerrors.Error)
	UpdateGitRepositoryById(id string, data challengemodels.UpdateGitRepositoryData) grerrors.Error

	StartGitRepositoryScanning(id string) (*grmodels.GitRepositoryScanResult, grerrors.Error)
}
