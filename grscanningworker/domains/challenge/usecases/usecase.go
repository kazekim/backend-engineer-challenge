package challengeusecases

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

type UseCase interface {
	DoGitRepositoryScanningForResultId(resultId string) grerrors.Error
}
