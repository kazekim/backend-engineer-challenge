package grcgitrepositorydb

import (
	grcgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

//DeleteGitRepositoryById delete git repository row by id
func (c *defaultClient) DeleteGitRepositoryById(id string) grerrors.Error {

	var dao grcgitrepositorydbdaos.GitRepository
	wq := "id = $1"
	vErr := c.db.DeleteExec(&dao, wq, id)
	if vErr != nil {
		return vErr
	}
	return nil
}
