package challengehandlers

import (
	challengeusecases "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/usecases"
	grcenv "github.com/kazekim/backend-engineer-challenge/grchallengeapi/env"
)

// Handler token handler
type defaultHandler struct {
	cfg *grcenv.Environment
	cu  challengeusecases.UseCase
}

// NewHandler new handler
func NewHandler(c *grcenv.Environment, cu challengeusecases.UseCase) Handler {
	return &defaultHandler{
		cfg: c,
		cu:  cu,
	}
}
