package challengehandlers

import begincontext "github.com/kazekim/backend-engineer-challenge/grlib/begin/context"

// Handler user interface
type Handler interface {
	FrontCreateRepository(c *begincontext.Context)
}
