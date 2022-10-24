package grmodels

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/bejson"
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grenums"
	"time"
)

var MockGitRepositoryScanResultWithDetail = GitRepositoryScanResultWithDetail{
	Id:             "xxxx",
	CreatedAt:      time.Now(),
	UpdatedAt:      time.Now(),
	RepositoryName: "name",
	RepositoryUrl:  "url",
	Status:         grenums.ScanStatusQueued,
	Findings:       bejson.JSON("test"),
	QueuedAt:       time.Now(),
}

type GitRepositoryScanResultWithDetail struct {
	Id             string             `json:"id"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
	RepositoryName string             `json:"repository_name"`
	RepositoryUrl  string             `json:"repository_url"`
	Status         grenums.ScanStatus `json:"status"`
	Findings       bejson.JSON        `json:"findings"`
	QueuedAt       time.Time          `json:"queued_at"`
	ScanningAt     *time.Time         `json:"scanning_at"`
	FinishedAt     *time.Time         `json:"finished_at"`
}

func ParseGitRepositoryScanResultWithDetailFromDao(dao *grgitrepositorydbdaos.GitRepositoryScanResultWithDetail) *GitRepositoryScanResultWithDetail {
	m := &GitRepositoryScanResultWithDetail{
		Id:             dao.Id,
		CreatedAt:      dao.CreatedAt,
		UpdatedAt:      dao.UpdatedAt,
		RepositoryName: dao.Name,
		RepositoryUrl:  dao.Url,
		Status:         dao.Status,
		Findings:       dao.Findings,
		QueuedAt:       dao.QueuedAt.Time,
	}

	if dao.ScanningAt.Valid {
		m.ScanningAt = &dao.ScanningAt.Time
	}
	if dao.FinishedAt.Valid {
		m.FinishedAt = &dao.FinishedAt.Time
	}

	return m
}

func ParseGitRepositoryScanResultsWithDetailFromDaos(daos *[]grgitrepositorydbdaos.GitRepositoryScanResultWithDetail) *[]GitRepositoryScanResultWithDetail {
	var ms []GitRepositoryScanResultWithDetail
	for _, dao := range *daos {
		m := ParseGitRepositoryScanResultWithDetailFromDao(&dao)
		ms = append(ms, *m)
	}
	return &ms
}
