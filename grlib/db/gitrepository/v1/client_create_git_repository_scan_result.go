package grgitrepositorydb

import (
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

// CreateGitRepositoryScanResult create git repository scan result row with params
func (c *defaultClient) CreateGitRepositoryScanResult(params *grgitrepositorydbdaos.GitRepositoryScanResult) (*grgitrepositorydbdaos.GitRepositoryScanResult, grerrors.Error) {

	q := "id, created_at, updated_at, repository_id, status, findings, queued_at, scanning_at, finished_at"
	v := ":id, :created_at, :updated_at, :repository_id, :status, :findings, :queued_at, :scanning_at, :finished_at"

	_, vErr := c.db.NamedExec(q, v, params)
	if vErr != nil {
		return nil, vErr
	}

	return params, nil
}
