package grcrouters

import "github.com/kazekim/backend-engineer-challenge/grlib/begin"

func (r *defaultRouter) registerFrontChallengeRouter(rg begin.RouterGroup) {

	rcg := rg.Group("/repository")
	{
		rcg.POST("", r.ch.FrontCreateRepository)
	}

}
