package grwmodels

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/befiles"
	grgitrepositorydb "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1"
	"github.com/kazekim/backend-engineer-challenge/grlib/grgitscanner"
	"github.com/kazekim/backend-engineer-challenge/grlib/grscankafka"
	grwenv "github.com/kazekim/backend-engineer-challenge/grscanningworker/env"
)

// Options option for new router
type Options struct {
	Environment           *grwenv.Environment
	FileClient            befiles.Client
	GitScannerClient      grgitscanner.Client
	GitRepositoryDBClient grgitrepositorydb.Client
	GitScannerMQClient    grscankafka.Client
}
