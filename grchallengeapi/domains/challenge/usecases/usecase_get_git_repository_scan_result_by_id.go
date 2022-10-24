package challengeusecases

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
)

//GetGitRepositoryScanResultById get git repository scan result data by reference id
func (u *defaultUseCase) GetGitRepositoryScanResultById(id string) (*grmodels.GitRepositoryScanResultWithDetail, grerrors.Error) {
	return u.cr.GetGitRepositoryScanResultById(id)
}
