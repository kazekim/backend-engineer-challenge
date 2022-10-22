package grcrouters

import (
	challengehandlers "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/handlers"
	grcmodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/models"
	"github.com/kazekim/backend-engineer-challenge/grlib/begin"
)

type Router interface {
	Gin() begin.Gin
	Start() error
}

type defaultRouter struct {
	server  begin.Gin
	ch      challengehandlers.Handler
	options *grcmodels.Options
}

// NewWithOptions new router with options
func NewWithOptions(options *grcmodels.Options) Router {
	if options.Environment == nil {
		panic("Not found Environment")
	}

	wf := begin.New()
	wf.RegisterAPIVersionGroup()

	r := &defaultRouter{
		server:  wf,
		options: options,
	}

	r.initHandlers()

	r.registerFrontRouter(wf.APIVersionGroup())

	return r
}

func (r *defaultRouter) Gin() begin.Gin {
	return r.server
}
