package begit

import "github.com/kazekim/backend-engineer-challenge/grlib/grerrors"

//Close clear all git temp data and end process
func (g *Git) Close() grerrors.Error {
	if g.fullPath == nil {
		return grerrors.ErrGitHasNotBeenLoadedYet
	}
	return g.c.FileClient().ForceDeleteDirectory(*g.fullPath)
}
