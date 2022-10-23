package challengehandlers

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	begincontext "github.com/kazekim/backend-engineer-challenge/grlib/begin/context"
)

//FrontGetGitRepositoryScanResultById get git repository scan result by id handler
func (h *defaultHandler) FrontGetGitRepositoryScanResultById(c *begincontext.Context) {

	var req challengemodels.FrontGetGitRepositoryScanResultByIdRequest
	vErr := c.BindAndValidate(&req)
	if vErr != nil {
		c.CreateResponseError(vErr)
		return
	}

	m, vErr := h.cu.GetGitRepositoryScanResultById(req.ResultId)
	if vErr != nil {
		c.CreateResponseError(vErr)
		return
	}

	resp := challengemodels.FrontGetGitRepositoryScanResultByIdResponse{
		Data: *m,
	}

	c.CreateResponseSuccess(resp)
}
