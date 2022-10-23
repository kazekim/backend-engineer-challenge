package challengehandlers

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	begincontext "github.com/kazekim/backend-engineer-challenge/grlib/begin/context"
)

//FrontGetGitRepositoryById get git repository by id handler
func (h *defaultHandler) FrontGetGitRepositoryById(c *begincontext.Context) {

	var req challengemodels.FrontGetGitRepositoryByIdRequest
	vErr := c.BindJSONAndValidate(&req)
	if vErr != nil {
		c.CreateResponseError(vErr)
		return
	}

	m, vErr := h.cu.GetGitRepositoryById(req.RepositoryId)
	if vErr != nil {
		c.CreateResponseError(vErr)
		return
	}
	
	resp := challengemodels.FrontGetGitRepositoryByIdResponse{
		Data: *m,
	}

	c.CreateResponseSuccess(resp)
}
