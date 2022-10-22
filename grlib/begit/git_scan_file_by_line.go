package begit

import (
	"bufio"
	"fmt"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"os"
)

//ScanFileByLine scan file and do action in rf func by line
func (g *defaultGit) ScanFileByLine(fileFullPath string, rf func(data string) grerrors.Error) grerrors.Error {

	if g.fullPath == nil {
		return grerrors.ErrGitHasNotBeenLoadedYet
	}

	path := fmt.Sprintf("%v/%v", *g.fullPath, fileFullPath)
	file, err := os.Open(path)
	if err != nil {
		return grerrors.NewDefaultError(err)
	}
	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)
	const maxCapacity int = 10000000
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		vErr := rf(scanner.Text())
		if vErr != nil {
			return vErr
		}
	}

	if err := scanner.Err(); err != nil {
		return grerrors.NewDefaultError(err)
	}
	return nil
}
