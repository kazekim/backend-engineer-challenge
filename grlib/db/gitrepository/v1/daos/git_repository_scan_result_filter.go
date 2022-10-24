package grgitrepositorydbdaos

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grenums"
)

type GitRepositoryScanResultFilter struct {
	Id           *string             `db:"id"`
	RepositoryId *string             `db:"repository_id"`
	Status       *grenums.ScanStatus `db:"status"`
}
