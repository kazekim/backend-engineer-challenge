package begit

import "github.com/kazekim/backend-engineer-challenge/grlib/grerrors"

type Git interface {
	Close() grerrors.Error

	Clone() grerrors.Error
	ListFiles() (*[]string, grerrors.Error)
	ScanFileByLine(fileFullPath string, rf func(data string) grerrors.Error) grerrors.Error
}

type defaultGit struct {
	c        Client
	name     string
	url      string
	path     *string
	fullPath *string
}
