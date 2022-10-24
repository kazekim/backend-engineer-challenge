package challengemodels

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grenums"
)

type GitRepositoryScanResultFilterData struct {
	Id           *string             `form:"id"`
	RepositoryId *string             `form:"repository_id"`
	Status       *grenums.ScanStatus `form:"status"`
}
