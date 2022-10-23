package grcgitrepositorydbdaos

import (
	grcenums "github.com/kazekim/backend-engineer-challenge/grchallengeapi/enums"
)

type GitRepositoryScanResultFilter struct {
	Id           *string              `db:"id"`
	RepositoryId *string              `db:"repository_id"`
	Status       *grcenums.ScanStatus `db:"status"`
}
