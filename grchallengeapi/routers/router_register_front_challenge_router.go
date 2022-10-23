package grcrouters

import (
	grcconst "github.com/kazekim/backend-engineer-challenge/grchallengeapi/const"
	"github.com/kazekim/backend-engineer-challenge/grlib/begin"
)

func (r *defaultRouter) registerFrontChallengeRouter(rg begin.RouterGroup) {

	rcg := rg.Group("/repository")
	{
		rcg.POST("", r.ch.FrontCreateGitRepository)
		rcg.GET("/:"+grcconst.ParamGitRepositoryId, r.ch.FrontGetGitRepositoryById)
		rcg.DELETE("/:"+grcconst.ParamGitRepositoryId, r.ch.FrontDeleteGitRepositoryById)
		rcg.PUT("/:"+grcconst.ParamGitRepositoryId, r.ch.FrontUpdateGitRepositoryById)
		rcg.POST("/list", r.ch.FrontListGitRepositories)

		rcg.POST("/:"+grcconst.ParamGitRepositoryId+"/scan/start", r.ch.FrontStartGitRepositoryScanning)
		rcg.POST("/:"+grcconst.ParamGitRepositoryId+"/scan_result/list", r.ch.FrontListGitRepositoryScanResultsByRepositoryId)

	}

	srcg := rg.Group("/scan_result")
	{
		srcg.POST("/list", r.ch.FrontListGitRepositoryScanResults)
		srcg.GET("/:"+grcconst.ParamGitRepositoryScanResultId, r.ch.FrontGetGitRepositoryScanResultById)
	}

}
