package grcgitrepositorydb

import (
	"database/sql"
	grcenums "github.com/kazekim/backend-engineer-challenge/grchallengeapi/enums"
	grcgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grchallengeapi/pkgs/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/bejson"
	"github.com/kazekim/backend-engineer-challenge/grlib/besqlx"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"time"
)

type UpdateGitRepositoryScanResultDB interface {
	Status(value grcenums.ScanStatus) UpdateGitRepositoryScanResultDB
	Findings(value interface{}) UpdateGitRepositoryScanResultDB
	QueuedAt(value time.Time) UpdateGitRepositoryScanResultDB
	ScanningAt(value time.Time) UpdateGitRepositoryScanResultDB
	FinishedAt(value time.Time) UpdateGitRepositoryScanResultDB

	Commit() grerrors.Error
}

type defaultUpdateGitRepositoryScanResultDB struct {
	helper *besqlx.UpdateHelper
}

//NewUpdateGitRepositoryScanResultDBById create git repository scan result update db helper
func NewUpdateGitRepositoryScanResultDBById(db besqlx.Client, id string) UpdateGitRepositoryScanResultDB {

	helper := db.NewUpdateHelper()
	helper.SetWhereParam("id", id)

	return &defaultUpdateGitRepositoryScanResultDB{
		helper: helper,
	}
}

//Commit do commit update git repository database data
func (u *defaultUpdateGitRepositoryScanResultDB) Commit() grerrors.Error {

	var m grcgitrepositorydbdaos.GitRepositoryScanResult

	vErr := u.helper.CommitUpdateQuery(m.TableName())
	if vErr != nil {
		return vErr
	}

	return nil
}

//Status set parameter to update status field
func (u *defaultUpdateGitRepositoryScanResultDB) Status(value grcenums.ScanStatus) UpdateGitRepositoryScanResultDB {
	u.helper.SetParam("status", value)
	return u
}

//Findings set parameter to update findings field
func (u *defaultUpdateGitRepositoryScanResultDB) Findings(value interface{}) UpdateGitRepositoryScanResultDB {

	b, _ := bejson.Marshal(value)
	u.helper.SetParam("findings", b)
	return u
}

//QueuedAt set parameter to update queued_at field
func (u *defaultUpdateGitRepositoryScanResultDB) QueuedAt(value time.Time) UpdateGitRepositoryScanResultDB {
	t := sql.NullTime{
		Time: value,
		Valid: true,
	}
	u.helper.SetParam("queued_at", t)
	return u
}

//ScanningAt set parameter to update scanning_at field
func (u *defaultUpdateGitRepositoryScanResultDB) ScanningAt(value time.Time) UpdateGitRepositoryScanResultDB {
	t := sql.NullTime{
		Time: value,
		Valid: true,
	}
	u.helper.SetParam("scanning_at", t)
	return u
}

//FinishedAt set parameter to update finished_at field
func (u *defaultUpdateGitRepositoryScanResultDB) FinishedAt(value time.Time) UpdateGitRepositoryScanResultDB {
	t := sql.NullTime{
		Time: value,
		Valid: true,
	}
	u.helper.SetParam("finished_at", t)
	return u
}
