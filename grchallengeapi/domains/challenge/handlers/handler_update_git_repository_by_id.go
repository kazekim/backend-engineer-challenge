package challengehandlers

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	begincontext "github.com/kazekim/backend-engineer-challenge/grlib/begin/context"
)

// FrontUpdateGitRepositoryById update git repository by id handler
func (h *defaultHandler) FrontUpdateGitRepositoryById(c *begincontext.Context) {

	var req challengemodels.FrontUpdateGitRepositoryByIdRequest
	vErr := c.BindAndValidate(&req)
	if vErr != nil {
		c.CreateResponseError(vErr)
		return
	}

	vErr = h.cu.UpdateGitRepositoryById(req.RepositoryId, req.UpdateGitRepositoryData)
	if vErr != nil {
		c.CreateResponseError(vErr)
		return
	}

	resp := challengemodels.FrontUpdateGitRepositoryByIdResponse{
		IsSuccess: true,
	}

	c.CreateResponseSuccess(resp)
}
