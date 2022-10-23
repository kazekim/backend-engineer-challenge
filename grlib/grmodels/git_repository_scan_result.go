package grmodels

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/bejson"
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grenums"
	"time"
)

type GitRepositoryScanResult struct {
	Id           string             `json:"id"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	RepositoryId string             `json:"repository_id"`
	Status       grenums.ScanStatus `json:"status"`
	Findings     bejson.JSON        `json:"findings"`
	QueuedAt     time.Time          `json:"queued_at"`
	ScanningAt   *time.Time         `json:"scanning_at"`
	FinishedAt   *time.Time         `json:"finished_at"`
}

func ParseGitRepositoryScanResultFromDao(dao *grgitrepositorydbdaos.GitRepositoryScanResult) *GitRepositoryScanResult {
	m := &GitRepositoryScanResult{
		Id:           dao.Id,
		CreatedAt:    dao.CreatedAt.Time,
		UpdatedAt:    dao.UpdatedAt.Time,
		RepositoryId: dao.RepositoryId,
		Status:       dao.Status,
		Findings:     dao.Findings,
		QueuedAt:     dao.QueuedAt.Time,
	}

	if dao.ScanningAt.Valid {
		m.ScanningAt = &dao.ScanningAt.Time
	}
	if dao.FinishedAt.Valid {
		m.FinishedAt = &dao.FinishedAt.Time
	}
	return m
}

func ParseGitRepositoryScanResultsFromDaos(daos *[]grgitrepositorydbdaos.GitRepositoryScanResult) *[]GitRepositoryScanResult {
	var ms []GitRepositoryScanResult
	for _, dao := range *daos {
		m := ParseGitRepositoryScanResultFromDao(&dao)
		ms = append(ms, *m)
	}
	return &ms
}
