package challengehandlers

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	grcmodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/models"
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

	m := &grcmodels.GitRepository{}

	resp := challengemodels.FrontGetGitRepositoryByIdResponse{
		GitRepository: *m,
	}

	c.CreateResponseSuccess(resp)
}
