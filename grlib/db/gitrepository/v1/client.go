package grcgitrepositorydb

import (
	grcgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

type Client interface {
	CreateGitRepository(params *grcgitrepositorydbdaos.GitRepository) (*grcgitrepositorydbdaos.GitRepository, grerrors.Error)
	DeleteGitRepositoryById(id string) grerrors.Error
	GetGitRepositoryById(id string) (*grcgitrepositorydbdaos.GitRepository, grerrors.Error)
	ListGitRepositories(filter grcgitrepositorydbdaos.GitRepositoryFilter, values ...int64) (*[]grcgitrepositorydbdaos.GitRepository, int64, grerrors.Error)
	UpdateGitRepositoryById(id string) UpdateGitRepositoryDB

	CreateGitRepositoryScanResult(params *grcgitrepositorydbdaos.GitRepositoryScanResult) (*grcgitrepositorydbdaos.GitRepositoryScanResult, grerrors.Error)
	DeleteGitRepositoryScanResultById(id string) grerrors.Error
	GetGitRepositoryScanResultById(id string) (*grcgitrepositorydbdaos.GitRepositoryScanResult, grerrors.Error)
	ListGitRepositoryScanResults(filter grcgitrepositorydbdaos.GitRepositoryScanResultFilter, values ...int64) (*[]grcgitrepositorydbdaos.GitRepositoryScanResultWithDetail, int64, grerrors.Error)
	UpdateGitRepositoryScanResultById(id string) UpdateGitRepositoryScanResultDB
}
