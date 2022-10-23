package challengerepositories

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

func (r *defaultRepository) DeleteGitRepositoryById(id string) grerrors.Error {

	vErr := r.grc.DeleteGitRepositoryById(id)
	if vErr != nil {
		return vErr
	}

	return nil
}
