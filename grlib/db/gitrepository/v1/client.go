package grgitrepositorydb

import (
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

type Client interface {
	CreateGitRepository(params *grgitrepositorydbdaos.GitRepository) (*grgitrepositorydbdaos.GitRepository, grerrors.Error)
	DeleteGitRepositoryById(id string) grerrors.Error
	GetGitRepositoryById(id string) (*grgitrepositorydbdaos.GitRepository, grerrors.Error)
	ListGitRepositories(filter grgitrepositorydbdaos.GitRepositoryFilter, values ...int64) (*[]grgitrepositorydbdaos.GitRepository, int64, grerrors.Error)
	UpdateGitRepositoryById(id string) UpdateGitRepositoryDB

	CreateGitRepositoryScanResult(params *grgitrepositorydbdaos.GitRepositoryScanResult) (*grgitrepositorydbdaos.GitRepositoryScanResult, grerrors.Error)
	DeleteGitRepositoryScanResultById(id string) grerrors.Error
	GetGitRepositoryScanResultById(id string) (*grgitrepositorydbdaos.GitRepositoryScanResult, grerrors.Error)
	ListGitRepositoryScanResults(filter grgitrepositorydbdaos.GitRepositoryScanResultFilter, values ...int64) (*[]grgitrepositorydbdaos.GitRepositoryScanResultWithDetail, int64, grerrors.Error)
	UpdateGitRepositoryScanResultById(id string) UpdateGitRepositoryScanResultDB
}
