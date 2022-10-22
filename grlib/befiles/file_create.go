package befiles

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"os"
)

func CreateFile(file string) grerrors.Error {
	f, err := os.Create(file)
	if err != nil {
		return grerrors.NewDefaultError(err)
	}

	defer f.Close()
	return nil
}
