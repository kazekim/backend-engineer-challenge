package challengerepositories

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
)

//GetGitRepositoryById get git repository entity with reference id by db client
func (r *defaultRepository) GetGitRepositoryById(id string) (*grmodels.GitRepository, grerrors.Error) {

	dao, vErr := r.grc.GetGitRepositoryById(id)
	if vErr != nil {
		return nil, vErr
	}

	m := grmodels.ParseGitRepositoryFromDao(dao)

	return m, nil
}
