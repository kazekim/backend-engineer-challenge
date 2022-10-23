package challengerepositories

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	grcmodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/models"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

type Repository interface {
	CreateGitRepository(data challengemodels.CreateGitRepositoryData) (*grcmodels.GitRepository, grerrors.Error)
	DeleteGitRepositoryById(id string) grerrors.Error
}
