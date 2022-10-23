package grcgitrepositorydbdaos

import (
	"database/sql"
	"github.com/kazekim/backend-engineer-challenge/grlib/bejson"
	"github.com/kazekim/backend-engineer-challenge/grlib/besqlx"
	"github.com/kazekim/backend-engineer-challenge/grlib/grenums"
)

type GitRepositoryScanResult struct {
	besqlx.IDModel
	RepositoryId string             `db:"repository_id"`
	Status       grenums.ScanStatus `db:"status"`
	Findings     bejson.JSON        `db:"findings"`
	QueuedAt     sql.NullTime       `db:"queued_at"`
	ScanningAt   sql.NullTime       `db:"scanning_at"`
	FinishedAt   sql.NullTime       `db:"finished_at"`
}

func (pf GitRepositoryScanResult) TableName() string {
	return "gr_git_repository_scan_results"
}
