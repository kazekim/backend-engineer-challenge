package challengehandlers

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	begincontext "github.com/kazekim/backend-engineer-challenge/grlib/begin/context"
)

// FrontListGitRepositories list git repositories handler
func (h *defaultHandler) FrontListGitRepositories(c *begincontext.Context) {

	var req challengemodels.FrontListGitRepositoriesRequest
	vErr := c.BindAndValidate(&req)
	if vErr != nil {
		c.CreateResponseError(vErr)
		return
	}

	ms, count, vErr := h.cu.ListGitRepositories(req.GitRepositoryFilterData, req.Page, req.Limit)
	if vErr != nil {
		c.CreateResponseError(vErr)
		return
	}

	page := req.Page
	if req.Page == 0 {
		page = 1
	}

	resp := challengemodels.FrontListGitRepositoriesResponse{
		Datas: *ms,
		Page:  page,
		Count: count,
	}

	c.CreateResponseSuccess(resp)
}
