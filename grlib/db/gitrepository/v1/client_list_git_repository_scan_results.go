package grcgitrepositorydb

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/besqlx"
	grcgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

//ListGitRepositoryScanResults list git repository scan result with filter
func (c *defaultClient) ListGitRepositoryScanResults(filter grcgitrepositorydbdaos.GitRepositoryScanResultFilter, values ...int64) (*[]grcgitrepositorydbdaos.GitRepositoryScanResultWithDetail, int64, grerrors.Error) {

	var daos []grcgitrepositorydbdaos.GitRepositoryScanResultWithDetail
	var dao grcgitrepositorydbdaos.GitRepositoryScanResultWithDetail
	wqb := besqlx.NewWhereQueryBuilder()
	if filter.Id != nil {
		wqb.Where("sr.id", *filter.Id, besqlx.WhereOperatorEqual)
	}
	if filter.RepositoryId != nil {
		wqb.Where("sr.repository_id", *filter.RepositoryId, besqlx.WhereOperatorEqual)
	}
	if filter.Status != nil {
		wqb.Where("sr.status", *filter.Status, besqlx.WhereOperatorEqual)
	}
	whereQuery, args := wqb.BuildQuery()

	orderByQuery := "sr.created_at desc"
	sb := besqlx.SelectQueryBuilder{
		SelectQuery:  "sr.id, gr.name as repository_name, gr.url as repository_url, sr.status, sr.findings, " +
			"sr.queued_at, sr.scanning_at, sr.finished_at",
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