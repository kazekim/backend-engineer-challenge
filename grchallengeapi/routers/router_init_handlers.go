package grcrouters

import (
	challengehandlers "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/handlers"
	challengerepositories "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/repositories"
	challengeusecases "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/usecases"
)

func (r *defaultRouter) initHandlers() {

	//Repositories
	cr, _ := challengerepositories.NewRepository(r.options.Environment)

	//UseCases
	cu := challengeusecases.NewUseCase(r.options.Environment, cr)

	//Handlers
	ch := challengehandlers.NewHandler(r.options.Environment, cu)

	//Assign to router
	r.ch = ch
}
