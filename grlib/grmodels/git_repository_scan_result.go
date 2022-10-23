package grmodels

import (
	"database/sql"
	"github.com/kazekim/backend-engineer-challenge/grlib/bejson"
	grcgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grenums"
	"time"
)

type GitRepositoryScanResult struct {
	Id           string              `json:"id"`
	CreatedAt    time.Time           `json:"created_at"`
	UpdatedAt    time.Time           `json:"updated_at"`
	RepositoryId string              `json:"repository_id"`
	Status       grenums.ScanStatus `json:"status"`
	Findings     bejson.JSON         `json:"findings"`
	QueuedAt     sql.NullTime        `json:"queued_at"`
	ScanningAt   sql.NullTime        `json:"scanning_at"`
	FinishedAt   sql.NullTime        `json:"finished_at"`
}

func ParseGitRepositoryScanResultFromDao(dao *grcgitrepositorydbdaos.GitRepositoryScanResult) *GitRepositoryScanResult {
	return &GitRepositoryScanResult{
		Id:           dao.Id,
		CreatedAt:    dao.CreatedAt.Time,
		UpdatedAt:    dao.UpdatedAt.Time,
		RepositoryId: dao.RepositoryId,
		Status:       dao.Status,
		Findings:     dao.Findings,
		QueuedAt:     dao.QueuedAt,
		ScanningAt:   dao.ScanningAt,
		FinishedAt:   dao.FinishedAt,
	}
}
