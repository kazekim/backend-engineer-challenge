package grgitrepositorydb

import (
	"fmt"
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

//GetGitRepositoryScanResultById get git repository scan result data by id
func (c *defaultClient) GetGitRepositoryScanResultById(id string) (*grgitrepositorydbdaos.GitRepositoryScanResultWithDetail, grerrors.Error) {

	var dao grgitrepositorydbdaos.GitRepositoryScanResultWithDetail
	q := fmt.Sprintf("select sr.id, sr.created_at, sr.updated_at, gr.name as repository_name, gr.url as repository_url, sr.status, sr.findings, "+
		"sr.queued_at, sr.scanning_at, sr.finished_at from %v where sr.id = $1", dao.TableName())
	vErr := c.db.Get(&dao, q, id)
	if vErr != nil {
		if vErr == grerrors.ErrDataNotFound {
			return nil, grerrors.NewError(grerrors.ErrCodeDataNotFound, "scan result id not found")
		} else {
			return nil, vErr
		}
	}

	return &dao, nil
}
