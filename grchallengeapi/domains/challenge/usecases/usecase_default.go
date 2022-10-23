package challengeusecases

import (
	challengerepositories "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/repositories"
	grcenv "github.com/kazekim/backend-engineer-challenge/grchallengeapi/env"
	"github.com/kazekim/backend-engineer-challenge/grlib/grscankafka"
)

type defaultUseCase struct {
	cfg *grcenv.Environment
	cr  challengerepositories.Repository
	grmqc grscankafka.Client
}

// NewUseCase new challenge use case
func NewUseCase(cfg *grcenv.Environment, cr challengerepositories.Repository, grmqc grscankafka.Client) UseCase {

	return &defaultUseCase{
		cfg: cfg,
		cr:  cr,
		grmqc: grmqc,
	}
}
