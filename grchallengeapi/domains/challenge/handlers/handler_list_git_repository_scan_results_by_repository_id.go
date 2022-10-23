package challengehandlers

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	begincontext "github.com/kazekim/backend-engineer-challenge/grlib/begin/context"
)

//FrontListGitRepositoryScanResultsByRepositoryId list git repository scan results by repository id handler
func (h *defaultHandler) FrontListGitRepositoryScanResultsByRepositoryId(c *begincontext.Context) {

	var req challengemodels.FrontListGitRepositoryScanResultsByRepositoryIdRequest
	vErr := c.BindAndValidate(&req)
	if vErr != nil {
		c.CreateResponseError(vErr)
		return
	}

	req.GitRepositoryScanResultFilterData.RepositoryId = &req.RepositoryId

	ms, count, vErr := h.cu.ListGitRepositoryScanResults(req.GitRepositoryScanResultFilterData, req.Page, req.Limit)
	if vErr != nil {
		c.CreateResponseError(vErr)
		return
	}

	page := req.Page
	if req.Page == 0 {
		page = 1
	}

	resp := challengemodels.FrontListGitRepositoryScanResultsResponse{
		Datas: *ms,
		Page:  page,
		Count: count,
	}

	c.CreateResponseSuccess(resp)
}
