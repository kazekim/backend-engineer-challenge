package grcrouters

import "github.com/kazekim/backend-engineer-challenge/grlib/begin"

func (r *defaultRouter) registerFrontRouter(rg begin.RouterGroup) {

	cg := rg.Group("/front")

	r.registerFrontChallengeRouter(cg)
}
