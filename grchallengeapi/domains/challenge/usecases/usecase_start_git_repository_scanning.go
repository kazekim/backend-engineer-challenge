package challengeusecases

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
)

//StartGitRepositoryScanning start git repository scanning by repository id
func (u *defaultUseCase) StartGitRepositoryScanning(id string) (*grmodels.GitRepositoryScanResult, grerrors.Error) {

	m, vErr := u.cr.StartGitRepositoryScanning(id)
	if vErr != nil {
		return nil, vErr
	}

	//TODO: add scan queue here

	return m, nil
}
