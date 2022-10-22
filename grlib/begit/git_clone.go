package begit

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/rs/xid"
	"os"
)

//Clone do clone git from repository
func (g *defaultGit) Clone() grerrors.Error {

	path := xid.New().String()
	fullPath := fmt.Sprintf("%v/%v", g.c.Config().RootPath, path)

	_, err := git.PlainClone(fullPath, false, &git.CloneOptions{
		URL:      g.url,
		Progress: os.Stdout,
	})
	if err != nil {
		return grerrors.NewDefaultError(err)
	}

	g.path = &path
	g.fullPath = &fullPath

	return nil
}
