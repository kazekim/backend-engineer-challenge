package challengeusecases

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
	"github.com/kazekim/backend-engineer-challenge/grlib/grscankafka"
)

//StartGitRepositoryScanning start git repository scanning by repository id
func (u *defaultUseCase) StartGitRepositoryScanning(id string) (*grmodels.GitRepositoryScanResult, grerrors.Error) {

	m, vErr := u.cr.StartGitRepositoryScanning(id)
	if vErr != nil {
		return nil, vErr
	}

	data := grscankafka.PublishStartGitRepositoryScanningMessageParams{
		ResultId:        m.Id,
		GitRepositoryId: m.RepositoryId,
	}
	vErr = u.grmqc.PublishStartGitRepositoryScanningMessage(u.cfg.GitScannerMQConfig.ServerId, data)
	if vErr != nil {
		return nil, vErr
	}

	return m, nil
}
