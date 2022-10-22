package befiles

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

//ForceDeleteDirectory force delete specified directory
func (c *defaultClient) ForceDeleteDirectory(name string) grerrors.Error {
	return ForceDeleteDirectory(name)
}
