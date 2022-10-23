package main

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/beenv"
	"github.com/kazekim/backend-engineer-challenge/grlib/befiles"
	"github.com/kazekim/backend-engineer-challenge/grlib/besqlx"
	grgitrepositorydb "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1"
	"github.com/kazekim/backend-engineer-challenge/grlib/grgitscanner"
	"github.com/kazekim/backend-engineer-challenge/grlib/grscankafka"
	grwenv "github.com/kazekim/backend-engineer-challenge/grscanningworker/env"
	grwmodels "github.com/kazekim/backend-engineer-challenge/grscanningworker/models"
	grwrouters "github.com/kazekim/backend-engineer-challenge/grscanningworker/routers"
)

func main() {

	var cfg grwenv.Environment
	err := beenv.LoadConfig(&cfg)
	if err != nil {
		panic(err)
	}

	fc, vErr := befiles.NewClient("temp")
	if vErr != nil {
		panic(vErr)
	}

	gsc := grgitscanner.NewClient(&cfg.GitConfig, fc)

	gskc := grscankafka.NewClient(&cfg.GitScannerMQConfig)

	dbc := besqlx.NewClient(&cfg.DatabaseConfig)
	err = dbc.Connect()
	if err != nil {
		panic(err)
	}

	grdbc := grgitrepositorydb.NewClient(dbc)

	options := &grwmodels.Options{
		Environment:           &cfg,
		FileClient:            fc,
		GitRepositoryDBClient: grdbc,
		GitScannerClient:  gsc,
		GitScannerMQClient: gskc,
	}

	// -----------
	// Router
	// -----------
	router := grwrouters.NewWithOptions(options)
	_ = router.Start()
}
