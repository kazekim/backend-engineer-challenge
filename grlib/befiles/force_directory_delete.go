package befiles

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"os"
)

//ForceDeleteDirectory force delete directory with path
func ForceDeleteDirectory(name string) grerrors.Error {

	if _, err := os.Stat(name); !os.IsNotExist(err) {

		err = os.RemoveAll(name)
		if err != nil {
			vErr := grerrors.ErrFileNotFound
			return vErr
		}
	}

	return nil
}
