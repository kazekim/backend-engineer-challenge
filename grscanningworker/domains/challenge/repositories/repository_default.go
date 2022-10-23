package challengerepositories

import (
	grgitrepositorydb "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	grwenv "github.com/kazekim/backend-engineer-challenge/grscanningworker/env"
)

type defaultRepository struct {
	cfg *grwenv.Environment
	grc grgitrepositorydb.Client
}

// NewRepository new challenge repository
func NewRepository(cfg *grwenv.Environment, grc grgitrepositorydb.Client) (Repository, grerrors.Error) {
	return &defaultRepository{
		cfg: cfg,
		grc: grc,
	}, nil
}
