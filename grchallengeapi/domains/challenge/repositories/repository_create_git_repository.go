package challengerepositories

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	grcmodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/models"
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

func (r *defaultRepository) CreateGitRepository(data challengemodels.CreateGitRepositoryData) (*grcmodels.GitRepository, grerrors.Error) {

	params := grgitrepositorydbdaos.GitRepository{
		Name: data.Name,
		Url: data.Url,
		IsActive: true,
	}

	dao, vErr := r.grc.CreateGitRepository(&params)
	if vErr != nil {
		return nil, vErr
	}

	m := grcmodels.ParseGitRepositoryFromDao(dao)

	return m, nil
}
