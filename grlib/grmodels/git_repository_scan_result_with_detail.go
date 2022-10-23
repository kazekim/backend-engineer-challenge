package grmodels

import (
	"database/sql"
	"github.com/kazekim/backend-engineer-challenge/grlib/bejson"
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grenums"
	"time"
)

type GitRepositoryScanResultWithDetail struct {
	Id             string             `json:"id"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
	RepositoryName string             `json:"repository_name"`
	RepositoryUrl  string             `json:"repository_url"`
	Status         grenums.ScanStatus `json:"status"`
	Findings       bejson.JSON        `json:"findings"`
	QueuedAt       sql.NullTime       `json:"queued_at"`
	ScanningAt     sql.NullTime       `json:"scanning_at"`
	FinishedAt     sql.NullTime       `json:"finished_at"`
}

func ParseGitRepositoryScanResultWithDetailFromDao(dao *grgitrepositorydbdaos.GitRepositoryScanResultWithDetail) *GitRepositoryScanResultWithDetail {
	return &GitRepositoryScanResultWithDetail{
		Id:             dao.Id,
		CreatedAt:      dao.CreatedAt,
		UpdatedAt:      dao.UpdatedAt,
		RepositoryName: dao.Name,
		RepositoryUrl:  dao.Url,
		Status:         dao.Status,
		Findings:       dao.Findings,
		QueuedAt:       dao.QueuedAt,
		ScanningAt:     dao.ScanningAt,
		FinishedAt:     dao.FinishedAt,
	}
}

func ParseGitRepositoryScanResultsWithDetailFromDaos(daos *[]grgitrepositorydbdaos.GitRepositoryScanResultWithDetail) *[]GitRepositoryScanResultWithDetail {
	var ms []GitRepositoryScanResultWithDetail
	for _, dao := range *daos {
		m := ParseGitRepositoryScanResultWithDetailFromDao(&dao)
		ms = append(ms, *m)
	}
	return &ms
}
