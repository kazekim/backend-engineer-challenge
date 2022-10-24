package grmodels

import (
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"time"
)

var MockGitRepository = GitRepository{
	Id:        "xxxxx",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	Name:      "test git",
	Url:       "test url",
	IsActive:  true,
}

type GitRepository struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	IsActive  bool      `json:"is_active"`
}

func ParseGitRepositoryFromDao(dao *grgitrepositorydbdaos.GitRepository) *GitRepository {
	return &GitRepository{
		Id:        dao.Id,
		CreatedAt: dao.CreatedAt.Time,
		UpdatedAt: dao.UpdatedAt.Time,
		Name:      dao.Name,
		Url:       dao.Url,
		IsActive:  dao.IsActive,
	}
}

func ParseGitRepositoriesFromDaos(daos *[]grgitrepositorydbdaos.GitRepository) *[]GitRepository {

	var ms []GitRepository
	for _, dao := range *daos {
		m := ParseGitRepositoryFromDao(&dao)
		ms = append(ms, *m)
	}
	return &ms
}
