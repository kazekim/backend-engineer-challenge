package befiles

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"os"
)

// CreateDirectory create directory with path
func CreateDirectory(path string) grerrors.Error {

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return nil
	}

	//Create a folder/directory at a full qualified path
	err := os.Mkdir(path, 0755)
	if err != nil {
		return grerrors.ErrCreateDirectoryFailed
	}
	return nil
}
