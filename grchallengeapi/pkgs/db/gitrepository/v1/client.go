package grcgitrepositorydb

import (
	grcgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grchallengeapi/pkgs/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

type Client interface {
	CreateGitRepository(params *grcgitrepositorydbdaos.GitRepository) (*grcgitrepositorydbdaos.GitRepository, grerrors.Error)
	DeleteGitRepositoryById(grId string) grerrors.Error
	GetGitRepositoryById(id string) (*grcgitrepositorydbdaos.GitRepository, grerrors.Error)
	ListGitRepositories(filter grcgitrepositorydbdaos.GitRepositoryFilter, values ...int64) (*[]grcgitrepositorydbdaos.GitRepository, int64, grerrors.Error)
}
