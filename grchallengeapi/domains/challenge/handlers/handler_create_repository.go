package challengehandlers

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	begincontext "github.com/kazekim/backend-engineer-challenge/grlib/begin/context"
)

//FrontCreateRepository create repository handler
func (h *defaultHandler) FrontCreateRepository(c *begincontext.Context) {

	var req challengemodels.FrontCreateRepositoryRequest
	vErr := c.BindJSONAndValidate(&req)
	if vErr != nil {
		c.CreateResponseError(vErr)
		return
	}

	id := "asdfg"

	resp := challengemodels.FrontCreateRepositoryResponse{
		RepositoryId: id,
	}

	c.CreateResponseSuccess(resp)
}
