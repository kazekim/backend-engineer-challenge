package challengeusecases

import "github.com/kazekim/backend-engineer-challenge/grlib/grerrors"

// DeleteGitRepositoryById delete git repository data by reference id
func (u *defaultUseCase) DeleteGitRepositoryById(id string) grerrors.Error {
	return u.cr.DeleteGitRepositoryById(id)
}
