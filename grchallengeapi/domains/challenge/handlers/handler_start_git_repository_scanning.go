package challengehandlers

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	begincontext "github.com/kazekim/backend-engineer-challenge/grlib/begin/context"
)

// FrontStartGitRepositoryScanning start git repository scanning handler
func (h *defaultHandler) FrontStartGitRepositoryScanning(c *begincontext.Context) {

	var req challengemodels.FrontStartGitRepositoryScanningRequest
	vErr := c.BindAndValidate(&req)
	if vErr != nil {
		c.CreateResponseError(vErr)
		return
	}

	m, vErr := h.cu.StartGitRepositoryScanning(req.RepositoryId)
	if vErr != nil {
		c.CreateResponseError(vErr)
		return
	}

	resp := challengemodels.FrontStartGitRepositoryScanningResponse{
		GitRepositoryScanResult: *m,
	}

	c.CreateResponseSuccess(resp)
}
