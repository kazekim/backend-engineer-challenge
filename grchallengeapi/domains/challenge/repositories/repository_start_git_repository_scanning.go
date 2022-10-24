package challengerepositories

import (
	"database/sql"
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grenums"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
	"time"
)

// StartGitRepositoryScanning create git repository scanning data by repository id with initial data
func (r *defaultRepository) StartGitRepositoryScanning(id string) (*grmodels.GitRepositoryScanResult, grerrors.Error) {

	_, vErr := r.grc.GetGitRepositoryById(id)
	if vErr != nil {
		return nil, vErr
	}

	params := grgitrepositorydbdaos.GitRepositoryScanResult{
		RepositoryId: id,
		Status:       grenums.ScanStatusQueued,
		QueuedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	dao, vErr := r.grc.CreateGitRepositoryScanResult(&params)
	if vErr != nil {
		return nil, vErr
	}

	m := grmodels.ParseGitRepositoryScanResultFromDao(dao)

	return m, nil
}
