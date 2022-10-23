package grcgitrepositorydbdaos

import (
	"database/sql"
	grcenums "github.com/kazekim/backend-engineer-challenge/grchallengeapi/enums"
	"github.com/kazekim/backend-engineer-challenge/grlib/bejson"
)

type GitRepositoryScanResultWithDetail struct {
	Id           *string              `db:"id"`
	Name     string `db:"repository_name"`
	Url      string `db:"repository_url"`
	Status       grcenums.ScanStatus `db:"status"`
	Findings     bejson.JSON         `db:"findings"`
	QueuedAt     sql.NullTime        `db:"queued_at"`
	ScanningAt   sql.NullTime        `db:"scanning_at"`
	FinishedAt   sql.NullTime        `db:"finished_at"`
}

func (pf GitRepositoryScanResultWithDetail) TableName() string {
	return "gr_git_repositories gr join gr_git_repository_scan_results sr on gr.id = sr.repository_id"
}