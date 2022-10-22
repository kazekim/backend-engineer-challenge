package challengerepositories

import (
	grcenv "github.com/kazekim/backend-engineer-challenge/grchallengeapi/env"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

type defaultRepository struct {
	cfg *grcenv.Environment
}

// NewRepository new challenge repository
func NewRepository(cfg *grcenv.Environment) (Repository, grerrors.Error) {
	return &defaultRepository{
		cfg: cfg,
	}, nil
}
