package grgitrepositorydb

import (
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

//DeleteGitRepositoryScanResultById delete git repository scan result row by id
func (c *defaultClient) DeleteGitRepositoryScanResultById(id string) grerrors.Error {

	var dao grgitrepositorydbdaos.GitRepositoryScanResult
	wq := "id = $1"
	vErr := c.db.DeleteExec(&dao, wq, id)
	if vErr != nil {
		return vErr
	}
	return nil
}
