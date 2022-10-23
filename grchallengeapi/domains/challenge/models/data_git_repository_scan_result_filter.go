package challengemodels

import (
	grcenums "github.com/kazekim/backend-engineer-challenge/grchallengeapi/enums"
)

type GitRepositoryScanResultFilterData struct {
	Id           *string              `json:"id"`
	RepositoryId *string              `json:"repository_id"`
	Status       *grcenums.ScanStatus `json:"status"`
}
