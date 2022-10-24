package challengerepositories

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
)

// CreateGitRepository parse input from api to params and create git repository with db client
func (r *defaultRepository) CreateGitRepository(data challengemodels.CreateGitRepositoryData) (*grmodels.GitRepository, grerrors.Error) {

	params := grgitrepositorydbdaos.GitRepository{
		Name:     data.Name,
		Url:      data.Url,
		IsActive: true,
	}

	dao, vErr := r.grc.CreateGitRepository(&params)
	if vErr != nil {
		return nil, vErr
	}

	m := grmodels.ParseGitRepositoryFromDao(dao)

	return m, nil
}
