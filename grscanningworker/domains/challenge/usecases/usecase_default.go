package challengeusecases

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grgitscanner"
	challengerepositories "github.com/kazekim/backend-engineer-challenge/grscanningworker/domains/challenge/repositories"
	grwenv "github.com/kazekim/backend-engineer-challenge/grscanningworker/env"
)

type defaultUseCase struct {
	cfg *grwenv.Environment
	cr  challengerepositories.Repository
	gsc grgitscanner.Client
}

// NewUseCase new challenge use case
func NewUseCase(cfg *grwenv.Environment, cr challengerepositories.Repository, gsc grgitscanner.Client) UseCase {

	return &defaultUseCase{
		cfg: cfg,
		cr:  cr,
		gsc: gsc,
	}
}
