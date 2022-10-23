package challengehandlers

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	begincontext "github.com/kazekim/backend-engineer-challenge/grlib/begin/context"
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
)

//FrontStartGitRepositoryScanning start git repository scanning handler
func (h *defaultHandler) FrontStartGitRepositoryScanning(c *begincontext.Context) {

	var req challengemodels.FrontStartGitRepositoryScanningRequest
	vErr := c.BindJSONAndValidate(&req)
	if vErr != nil {
		c.CreateResponseError(vErr)
		return
	}

	m := &grmodels.GitRepositoryScanResult{}

	resp := challengemodels.FrontStartGitRepositoryScanningResponse{
		Data: *m,
	}

	c.CreateResponseSuccess(resp)
}
