package challengehandlers

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	grcmodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/models"
	begincontext "github.com/kazekim/backend-engineer-challenge/grlib/begin/context"
)

//FrontListGitRepositories list git repositories handler
func (h *defaultHandler) FrontListGitRepositories(c *begincontext.Context) {

	var req challengemodels.FrontListGitRepositoriesByIdRequest
	vErr := c.BindJSONAndValidate(&req)
	if vErr != nil {
		c.CreateResponseError(vErr)
		return
	}

	m := &[]grcmodels.GitRepository{}

	resp := challengemodels.FrontListGitRepositoriesResponse{
		Datas: *m,
		Page: 1,
		Count: 2,
	}

	c.CreateResponseSuccess(resp)
}
