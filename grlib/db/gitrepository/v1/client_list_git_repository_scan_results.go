package grgitrepositorydb

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/besqlx"
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

// ListGitRepositoryScanResults list git repository scan result with filter
func (c *defaultClient) ListGitRepositoryScanResults(filter grgitrepositorydbdaos.GitRepositoryScanResultFilter, values ...int64) (*[]grgitrepositorydbdaos.GitRepositoryScanResultWithDetail, int64, grerrors.Error) {

	var daos []grgitrepositorydbdaos.GitRepositoryScanResultWithDetail
	var dao grgitrepositorydbdaos.GitRepositoryScanResultWithDetail
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
		SelectQuery: "sr.id, sr.created_at, sr.updated_at, gr.name as repository_name, gr.url as repository_url, sr.status, sr.findings, " +
			"sr.queued_at, sr.scanning_at, sr.finished_at",
		WhereQuery:   whereQuery,
		OrderByQuery: &orderByQuery,
		Args:         args,
	}

	count, vErr := c.db.SelectWithCount(&dao, &daos, sb, values...)
	if vErr != nil {
		return nil, 0, vErr
	}

	return &daos, count, nil
}
