package grcgitrepositorydb

import (
	grcgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grchallengeapi/pkgs/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

func (c *defaultClient) DeleteGitRepositoryById(grId string) grerrors.Error {

	var dao grcgitrepositorydbdaos.GitRepository
	wq := "id = $1"
	vErr := c.db.DeleteExec(&dao, wq, grId)
	if vErr != nil {
		return vErr
	}
	return nil
}