package challengeusecases

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
)

// CreateGitRepository create git repository from input data
func (u *defaultUseCase) CreateGitRepository(data challengemodels.CreateGitRepositoryData) (*grmodels.GitRepository, grerrors.Error) {

	return u.cr.CreateGitRepository(data)
}
