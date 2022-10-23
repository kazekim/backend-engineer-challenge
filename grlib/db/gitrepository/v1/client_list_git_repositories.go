package grgitrepositorydb

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/besqlx"
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

//ListGitRepositories list git repositories with filter
func (c *defaultClient) ListGitRepositories(filter grgitrepositorydbdaos.GitRepositoryFilter, values ...int64) (*[]grgitrepositorydbdaos.GitRepository, int64, grerrors.Error) {

	var daos []grgitrepositorydbdaos.GitRepository
	var dao grgitrepositorydbdaos.GitRepository
	wqb := besqlx.NewWhereQueryBuilder()
	if filter.Id != nil {
		wqb.Where("id", *filter.Id, besqlx.WhereOperatorEqual)
	}
	if filter.Name != nil {
		wqb.Where("name", *filter.Name, besqlx.WhereOperatorLike)
	}
	if filter.IsActive != nil {
		wqb.Where("is_active", *filter.IsActive, besqlx.WhereOperatorEqual)
	}
	whereQuery, args := wqb.BuildQuery()

	orderByQuery := "created_at"
	sb := besqlx.SelectQueryBuilder{
		SelectQuery:  "*",
		WhereQuery:   whereQuery,
		OrderByQuery: &orderByQuery,
		Args:         args,
	}

	count, vErr := c.db.SelectWithCount(&dao, &daos, sb)
	if vErr != nil {
		return nil, 0, vErr
	}

	return &daos, count, nil
}
