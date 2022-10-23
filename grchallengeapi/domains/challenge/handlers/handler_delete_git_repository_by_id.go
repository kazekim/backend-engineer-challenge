package challengehandlers

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	begincontext "github.com/kazekim/backend-engineer-challenge/grlib/begin/context"
)

//FrontDeleteGitRepositoryById delete git repository by id handler
func (h *defaultHandler) FrontDeleteGitRepositoryById(c *begincontext.Context) {

	var req challengemodels.FrontDeleteGitRepositoryByIdRequest
	vErr := c.BindJSONAndValidate(&req)
	if vErr != nil {
		c.CreateResponseError(vErr)
		return
	}

	vErr = h.cu.DeleteGitRepositoryById(req.RepositoryId)
	if vErr != nil {
		c.CreateResponseError(vErr)
		return
	}

	resp := challengemodels.FrontDeleteGitRepositoryByIdResponse{
		IsSuccess: true,
	}

	c.CreateResponseSuccess(resp)
}
