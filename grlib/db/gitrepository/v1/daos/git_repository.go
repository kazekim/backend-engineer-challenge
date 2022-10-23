package grgitrepositorydbdaos

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/besqlx"
)

type GitRepository struct {
	besqlx.IDModel
	Name     string `db:"name"`
	Url      string `db:"url"`
	IsActive bool   `db:"is_active"`
}

func (pf GitRepository) TableName() string {
	return "gr_git_repositories"
}
