package main

import (
	grcenv "github.com/kazekim/backend-engineer-challenge/grchallengeapi/env"
	grcmodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/models"
	grcrouters "github.com/kazekim/backend-engineer-challenge/grchallengeapi/routers"
	"github.com/kazekim/backend-engineer-challenge/grlib/beenv"
	"github.com/kazekim/backend-engineer-challenge/grlib/befiles"
	"github.com/kazekim/backend-engineer-challenge/grlib/besqlx"
	grgitrepositorydb "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1"
	"github.com/kazekim/backend-engineer-challenge/grlib/grgitscanner"
)

func main() {

	var cfg grcenv.Environment
	err := beenv.LoadConfig(&cfg)
	if err != nil {
		panic(err)
	}

	fc, vErr := befiles.NewClient("temp")
	if vErr != nil {
		panic(vErr)
	}

	gsc := grgitscanner.NewClient(&cfg.GitConfig, fc)

	dbc := besqlx.NewClient(&cfg.DatabaseConfig)
	err = dbc.Connect()
	if err != nil {
		panic(err)
	}

	grdbc := grgitrepositorydb.NewClient(dbc)

	options := &grcmodels.Options{
		Environment:           &cfg,
		FileClient:            fc,
		GitScannerClient:      gsc,
		GitRepositoryDBClient: grdbc,
	}

	// -----------
	// Router
	// -----------
	router := grcrouters.NewWithOptions(options)
	_ = router.Start()
}
