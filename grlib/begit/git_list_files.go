package begit

import (
	"fmt"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"io/ioutil"
	"strings"
)

func (g *Git) ListFiles() (*[]string, grerrors.Error) {

	if g.fullPath == nil {
		return nil, grerrors.ErrGitHasNotBeenLoadedYet
	}

	files, vErr := g.listFilesInDir(*g.fullPath)
	if vErr != nil {
		return nil, vErr
	}

	return files, nil
}

func (g *Git) listFilesInDir(path string) (*[]string, grerrors.Error) {

	fs, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, grerrors.NewDefaultError(err)
	}

	var files []string
	for _, file := range fs {
		if !file.IsDir() {
			p := ""
			if path != *g.fullPath {
				p = strings.Replace(path, (*g.fullPath) + "/", "", -1)
				p = fmt.Sprintf("%v/", p)
			}
			fileFullPath := fmt.Sprintf("%v%v", p, file.Name())
			files = append(files, fileFullPath)
		}else{
			absPath := fmt.Sprintf("%v/%v", path, file.Name())
			fds, vErr := g.listFilesInDir(absPath)
			if vErr != nil {
				return nil, grerrors.NewDefaultError(err)
			}
			files = append(files, (*fds)...)
		}
	}
	return &files, nil
}