package challengeusecases

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

//UpdateGitRepositoryById update git repository data by reference id with update data
func (u *defaultUseCase) UpdateGitRepositoryById(id string, data challengemodels.UpdateGitRepositoryData) grerrors.Error {
	return u.cr.UpdateGitRepositoryById(id, data)
}
