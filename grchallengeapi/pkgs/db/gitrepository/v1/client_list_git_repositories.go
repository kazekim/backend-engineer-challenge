package grcgitrepositorydb

import (
	grcgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grchallengeapi/pkgs/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/besqlx"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

//ListGitRepositories list git repositories with filter
func (c *defaultClient) ListGitRepositories(filter grcgitrepositorydbdaos.GitRepositoryFilter, values ...int64) (*[]grcgitrepositorydbdaos.GitRepository, int64, grerrors.Error) {

	var daos []grcgitrepositorydbdaos.GitRepository
	var dao grcgitrepositorydbdaos.GitRepository
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
		SelectQuery: "*",
		WhereQuery: whereQuery,
		OrderByQuery: &orderByQuery,
		Args: args,
	}

	count, vErr := c.db.Count(&dao, whereQuery, args)
	if vErr != nil {
		return nil, 0, vErr
	}

	if len(values) == 2 {
		sb.Page = &values[0]
		sb.Limit = &values[1]
	}

	vErr = c.db.Select(&dao, &daos, sb)
	if vErr != nil {
		return nil, 0, vErr
	}

	return &daos, count, nil
}
