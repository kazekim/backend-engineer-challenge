package grgitrepositorydb

import (
	"fmt"
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

//GetGitRepositoryScanResultById get git repository scan result data by id
func (c *defaultClient) GetGitRepositoryScanResultById(id string) (*grgitrepositorydbdaos.GitRepositoryScanResultWithDetail, grerrors.Error) {

	var dao grgitrepositorydbdaos.GitRepositoryScanResultWithDetail
	q := fmt.Sprintf(`select * from %v where sr.id = $1`, dao.TableName())
	err := c.db.Get(&dao, q, id)
	if err != nil {
		return nil, grerrors.NewDatabaseError(err)
	}

	return &dao, nil
}
