package grcmodels

import (
	grcgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grchallengeapi/pkgs/db/gitrepository/v1/daos"
	"time"
)

type GitRepository struct {
	Id string `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	IsActive bool   `json:"is_active"`
}

func ParseGitRepositoryFromDao(dao *grcgitrepositorydbdaos.GitRepository) *GitRepository {
	return &GitRepository{
		Id: dao.Id,
		CreatedAt: dao.CreatedAt.Time,
		UpdatedAt: dao.UpdatedAt.Time,
		Name: dao.Name,
		Url: dao.Url,
		IsActive: dao.IsActive,
	}
}