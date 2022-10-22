package main

import (
	grcenv "github.com/kazekim/backend-engineer-challenge/grchallengeapi/env"
	grcmodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/models"
	"github.com/kazekim/backend-engineer-challenge/grchallengeapi/pkgs/grcgitscanner"
	grcrouters "github.com/kazekim/backend-engineer-challenge/grchallengeapi/routers"
	"github.com/kazekim/backend-engineer-challenge/grlib/beenv"
	"github.com/kazekim/backend-engineer-challenge/grlib/befiles"
	"github.com/kazekim/backend-engineer-challenge/grlib/besqlx"
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

	gsc := grcgitscanner.NewClient(&cfg.GitConfig, fc)

	dbc := besqlx.NewClient(&cfg.DatabaseConfig)
	err = dbc.Connect()
	if err != nil {
		panic(err)
	}

	options := &grcmodels.Options{
		Environment:      &cfg,
		FileClient:       fc,
		GitScannerClient: gsc,
	}

	// -----------
	// Router
	// -----------
	router := grcrouters.NewWithOptions(options)
	_ = router.Start()
}
