package grcgitrepositorydb

import (
	"fmt"
	grcgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

//GetGitRepositoryById get git repository data by id
func (c *defaultClient) GetGitRepositoryById(id string) (*grcgitrepositorydbdaos.GitRepository, grerrors.Error) {

	var dao grcgitrepositorydbdaos.GitRepository
	q := fmt.Sprintf(`select * from %v where id = $1`, dao.TableName())
	err := c.db.Get(&dao, q, id)
	if err != nil {
		return nil, grerrors.NewDatabaseError(err)
	}

	return &dao, nil
}
