package grcrouters

import (
	challengehandlers "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/handlers"
	challengerepositories "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/repositories"
	challengeusecases "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/usecases"
)

func (r *defaultRouter) initHandlers() {

	//Repositories
	cr, _ := challengerepositories.NewRepository(r.options.Environment, r.options.GitRepositoryDBClient)

	//UseCases
	cu := challengeusecases.NewUseCase(r.options.Environment, cr, r.options.GitScannerMQClient)

	//Handlers
	ch := challengehandlers.NewHandler(r.options.Environment, cu)

	//Assign to router
	r.ch = ch
}
