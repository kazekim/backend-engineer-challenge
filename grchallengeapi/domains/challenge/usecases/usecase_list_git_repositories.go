package challengeusecases

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
)

// ListGitRepositories list git repository datas by filter and page, limit data
func (u *defaultUseCase) ListGitRepositories(filter challengemodels.GitRepositoryFilterData, values ...int64) (*[]grmodels.GitRepository, int64, grerrors.Error) {
	return u.cr.ListGitRepositories(filter, values...)
}
