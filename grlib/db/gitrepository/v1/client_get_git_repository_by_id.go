package grgitrepositorydb

import (
	"fmt"
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

//GetGitRepositoryById get git repository data by id
func (c *defaultClient) GetGitRepositoryById(id string) (*grgitrepositorydbdaos.GitRepository, grerrors.Error) {

	var dao grgitrepositorydbdaos.GitRepository
	q := fmt.Sprintf(`select * from %v where id = $1`, dao.TableName())
	vErr := c.db.Get(&dao, q, id)
	if vErr != nil {
		if vErr == grerrors.ErrDataNotFound {
			return nil, grerrors.NewDatabaseErrorWithMessage("git repository id not found")
		} else {
			return nil, vErr
		}
	}

	return &dao, nil
}
