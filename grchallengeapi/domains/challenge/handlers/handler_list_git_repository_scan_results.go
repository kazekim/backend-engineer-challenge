package challengehandlers

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	grcmodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/models"
	begincontext "github.com/kazekim/backend-engineer-challenge/grlib/begin/context"
)

//FrontListGitRepositoryScanResults list git repository scan results handler
func (h *defaultHandler) FrontListGitRepositoryScanResults(c *begincontext.Context) {

	var req challengemodels.FrontListGitRepositoryScanResultsRequest
	vErr := c.BindJSONAndValidate(&req)
	if vErr != nil {
		c.CreateResponseError(vErr)
		return
	}

	m := &[]grcmodels.GitRepositoryScanResult{}
	page := int64(1)
	count := int64(2)

	resp := challengemodels.FrontListGitRepositoryScanResultsResponse{
		Datas: *m,
		Page: page,
		Count: count,
	}

	c.CreateResponseSuccess(resp)
}
