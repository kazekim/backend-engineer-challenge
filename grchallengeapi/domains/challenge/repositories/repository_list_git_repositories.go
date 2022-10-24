package challengerepositories

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
)

// ListGitRepositories parse input filter data to params and list git repository entities by db client
func (r *defaultRepository) ListGitRepositories(filter challengemodels.GitRepositoryFilterData, values ...int64) (*[]grmodels.GitRepository, int64, grerrors.Error) {

	fParams := grgitrepositorydbdaos.GitRepositoryFilter{
		Id:       filter.Id,
		Name:     filter.Name,
		IsActive: filter.IsActive,
	}

	daos, count, vErr := r.grc.ListGitRepositories(fParams, values...)
	if vErr != nil {
		return nil, 0, vErr
	}

	ms := grmodels.ParseGitRepositoriesFromDaos(daos)

	return ms, count, nil
}
