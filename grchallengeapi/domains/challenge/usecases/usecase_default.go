package challengeusecases

import (
	challengerepositories "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/repositories"
	grcenv "github.com/kazekim/backend-engineer-challenge/grchallengeapi/env"
)

type defaultUseCase struct {
	cfg *grcenv.Environment
	cr  challengerepositories.Repository
}

// NewUseCase new challenge use case
func NewUseCase(cfg *grcenv.Environment, cr challengerepositories.Repository) UseCase {

	return &defaultUseCase{
		cfg: cfg,
		cr:  cr,
	}
}
