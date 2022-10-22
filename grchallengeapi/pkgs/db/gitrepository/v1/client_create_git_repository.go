package grcgitrepositorydb

import (
	grcgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grchallengeapi/pkgs/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

func (c *defaultClient) CreateGitRepository(params *grcgitrepositorydbdaos.GitRepository) (*grcgitrepositorydbdaos.GitRepository, grerrors.Error) {

	q := "id, created_at, updated_at, name, url, is_active"
	v := ":id, :created_at, :updated_at, :name, :url, :is_active"

	_, vErr := c.db.NamedExec(q, v, params)
	if vErr != nil {
		return nil, vErr
	}

	return params, nil
}
