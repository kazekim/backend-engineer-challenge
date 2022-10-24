package challengerepositories

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
)

// ListGitRepositoryScanResults parse input filter data to params and list git repository scan result entities by db client
func (r *defaultRepository) ListGitRepositoryScanResults(filter challengemodels.GitRepositoryScanResultFilterData, values ...int64) (*[]grmodels.GitRepositoryScanResultWithDetail, int64, grerrors.Error) {

	fParams := grgitrepositorydbdaos.GitRepositoryScanResultFilter{
		Id:           filter.Id,
		RepositoryId: filter.RepositoryId,
		Status:       filter.Status,
	}

	daos, count, vErr := r.grc.ListGitRepositoryScanResults(fParams, values...)
	if vErr != nil {
		return nil, 0, vErr
	}

	ms := grmodels.ParseGitRepositoryScanResultsWithDetailFromDaos(daos)

	return ms, count, nil
}
