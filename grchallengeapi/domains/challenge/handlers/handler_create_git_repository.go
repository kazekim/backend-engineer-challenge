package challengehandlers

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	begincontext "github.com/kazekim/backend-engineer-challenge/grlib/begin/context"
)

//FrontCreateGitRepository create git repository handler
func (h *defaultHandler) FrontCreateGitRepository(c *begincontext.Context) {

	var req challengemodels.FrontCreateGitRepositoryRequest
	vErr := c.BindJSONAndValidate(&req)
	if vErr != nil {
		c.CreateResponseError(vErr)
		return
	}

	id := "asdfg"

	resp := challengemodels.FrontCreateGitRepositoryResponse{
		RepositoryId: id,
	}

	c.CreateResponseSuccess(resp)
}
