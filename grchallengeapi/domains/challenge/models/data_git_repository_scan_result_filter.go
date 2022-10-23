package challengemodels

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grenums"
)

type GitRepositoryScanResultFilterData struct {
	Id           *string              `json:"id"`
	RepositoryId *string              `json:"repository_id"`
	Status       *grenums.ScanStatus `json:"status"`
}
