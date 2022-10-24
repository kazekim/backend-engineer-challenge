package grgitrepositorydb

import (
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

// CreateGitRepository create git repository row with params
func (c *defaultClient) CreateGitRepository(params *grgitrepositorydbdaos.GitRepository) (*grgitrepositorydbdaos.GitRepository, grerrors.Error) {

	q := "id, created_at, updated_at, name, url, is_active"
	v := ":id, :created_at, :updated_at, :name, :url, :is_active"

	_, vErr := c.db.NamedExec(q, v, params)
	if vErr != nil {
		return nil, vErr
	}

	return params, nil
}
