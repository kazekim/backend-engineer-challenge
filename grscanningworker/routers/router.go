package grwrouters

import (
	challengeusecases "github.com/kazekim/backend-engineer-challenge/grscanningworker/domains/challenge/usecases"
	grwmodels "github.com/kazekim/backend-engineer-challenge/grscanningworker/models"
)

type Router interface {
	Start() error
}

type defaultRouter struct {
	cu      challengeusecases.UseCase
	options *grwmodels.Options
}

// NewWithOptions new router with options
func NewWithOptions(options *grwmodels.Options) Router {
	if options.Environment == nil {
		panic("Not found Environment")
	}

	r := &defaultRouter{
		options: options,
	}

	r.initHandlers()
	r.registerGitScannerWorker()

	return r
}
