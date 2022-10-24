package challengerepositories

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

// DeleteGitRepositoryScanResultById delete git repository scan result by result id with db client
func (r *defaultRepository) DeleteGitRepositoryScanResultById(id string) grerrors.Error {

	vErr := r.grc.DeleteGitRepositoryScanResultById(id)
	if vErr != nil {
		return vErr
	}

	return nil
}
