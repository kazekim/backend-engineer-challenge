package grgitrepositorydb

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/besqlx"
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

type UpdateGitRepositoryDB interface {
	Name(value string) UpdateGitRepositoryDB
	Url(value string) UpdateGitRepositoryDB
	IsActive(value bool) UpdateGitRepositoryDB

	Commit() grerrors.Error
}

type defaultUpdateGitRepositoryDB struct {
	helper *besqlx.UpdateHelper
}

//NewUpdateGitRepositoryDBById create git repository update db helper
func NewUpdateGitRepositoryDBById(db besqlx.Client, id string) UpdateGitRepositoryDB {

	helper := db.NewUpdateHelper()
	helper.SetWhereParam("id", id)

	return &defaultUpdateGitRepositoryDB{
		helper: helper,
	}
}

//Commit do commit update git repository database data
func (u *defaultUpdateGitRepositoryDB) Commit() grerrors.Error {

	var m grgitrepositorydbdaos.GitRepository
	vErr := u.helper.CommitUpdateQuery(&m)
	if vErr != nil {
		return vErr
	}

	return nil
}

//Name set parameter to update name field
func (u *defaultUpdateGitRepositoryDB) Name(value string) UpdateGitRepositoryDB {
	u.helper.SetParam("name", value)
	return u
}

//Url set parameter to update url field
func (u *defaultUpdateGitRepositoryDB) Url(value string) UpdateGitRepositoryDB {
	u.helper.SetParam("url", value)
	return u
}

//IsActive set parameter to update is_active field
func (u *defaultUpdateGitRepositoryDB) IsActive(value bool) UpdateGitRepositoryDB {
	u.helper.SetParam("is_active", value)
	return u
}
