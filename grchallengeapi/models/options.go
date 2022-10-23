package grcmodels

import (
	grcenv "github.com/kazekim/backend-engineer-challenge/grchallengeapi/env"
	"github.com/kazekim/backend-engineer-challenge/grlib/befiles"
	grgitrepositorydb "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1"
	"github.com/kazekim/backend-engineer-challenge/grlib/grgitscanner"
)

// Options option for new router
type Options struct {
	Environment           *grcenv.Environment
	FileClient            befiles.Client
	GitScannerClient      grgitscanner.Client
	GitRepositoryDBClient grgitrepositorydb.Client
}
