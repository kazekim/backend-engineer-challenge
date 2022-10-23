package grgitrepositorydbdaos

import (
	"database/sql"
	"github.com/kazekim/backend-engineer-challenge/grlib/bejson"
	"github.com/kazekim/backend-engineer-challenge/grlib/grenums"
	"time"
)

type GitRepositoryScanResultWithDetail struct {
	Id         string             `db:"id"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
	Name       string             `db:"repository_name"`
	Url        string             `db:"repository_url"`
	Status     grenums.ScanStatus `db:"status"`
	Findings   bejson.JSON        `db:"findings"`
	QueuedAt   sql.NullTime       `db:"queued_at"`
	ScanningAt sql.NullTime       `db:"scanning_at"`
	FinishedAt sql.NullTime       `db:"finished_at"`
}

func (pf GitRepositoryScanResultWithDetail) TableName() string {
	return "gr_git_repositories gr join gr_git_repository_scan_results sr on gr.id = sr.repository_id"
}
