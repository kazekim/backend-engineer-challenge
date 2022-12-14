package challengerepositories

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
)

// GetGitRepositoryScanResultById get git repository scan result entity with reference id by db client
func (r *defaultRepository) GetGitRepositoryScanResultById(id string) (*grmodels.GitRepositoryScanResultWithDetail, grerrors.Error) {

	dao, vErr := r.grc.GetGitRepositoryScanResultById(id)
	if vErr != nil {
		return nil, vErr
	}

	m := grmodels.ParseGitRepositoryScanResultWithDetailFromDao(dao)

	return m, nil
}
