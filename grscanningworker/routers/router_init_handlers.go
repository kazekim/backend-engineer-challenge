package grwrouters

import (
	challengerepositories "github.com/kazekim/backend-engineer-challenge/grscanningworker/domains/challenge/repositories"
	challengeusecases "github.com/kazekim/backend-engineer-challenge/grscanningworker/domains/challenge/usecases"
)

func (r *defaultRouter) initHandlers() {

	//Repositories
	cr, _ := challengerepositories.NewRepository(r.options.Environment, r.options.GitRepositoryDBClient)

	//UseCases
	cu := challengeusecases.NewUseCase(r.options.Environment, cr, r.options.GitScannerClient)

	//Assign to router
	r.cu = cu
}
