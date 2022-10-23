package challengeusecases

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
)

//ListGitRepositoryScanResults list git repository scan result data by filter and page, limit data
func (u *defaultUseCase) ListGitRepositoryScanResults(filter challengemodels.GitRepositoryScanResultFilterData, values ...int64) (*[]grmodels.GitRepositoryScanResultWithDetail, int64, grerrors.Error) {
	return u.cr.ListGitRepositoryScanResults(filter, values...)
}
