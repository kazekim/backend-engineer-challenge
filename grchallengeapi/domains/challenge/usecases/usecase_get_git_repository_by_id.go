package challengeusecases

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
)

//GetGitRepositoryById get git repository data by reference id
func (u *defaultUseCase) GetGitRepositoryById(id string) (*grmodels.GitRepository, grerrors.Error) {
	return u.cr.GetGitRepositoryById(id)
}
