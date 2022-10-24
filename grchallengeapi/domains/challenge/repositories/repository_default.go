package challengerepositories

import (
	grcenv "github.com/kazekim/backend-engineer-challenge/grchallengeapi/env"
	grgitrepositorydb "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

type defaultRepository struct {
	cfg *grcenv.Environment
	grc grgitrepositorydb.Client
}

// NewRepository new challenge repository
func NewRepository(cfg *grcenv.Environment, grc grgitrepositorydb.Client) (Repository, grerrors.Error) {
	return &defaultRepository{
		cfg: cfg,
		grc: grc,
	}, nil
}
