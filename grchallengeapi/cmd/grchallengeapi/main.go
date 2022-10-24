package main

import (
	grcenv "github.com/kazekim/backend-engineer-challenge/grchallengeapi/env"
	grcmodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/models"
	grcrouters "github.com/kazekim/backend-engineer-challenge/grchallengeapi/routers"
	"github.com/kazekim/backend-engineer-challenge/grlib/beenv"
	"github.com/kazekim/backend-engineer-challenge/grlib/befiles"
	"github.com/kazekim/backend-engineer-challenge/grlib/besqlx"
	grgitrepositorydb "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1"
	"github.com/kazekim/backend-engineer-challenge/grlib/grscankafka"
	"log"
)

func main() {

	var cfg grcenv.Environment
	err := beenv.LoadConfig(&cfg)
	if err != nil {
		log.Println(err)
		return
	}

	fc, vErr := befiles.NewClient("temp")
	if vErr != nil {
		log.Println(vErr)
		return
	}

	dbc := besqlx.NewClient(&cfg.DatabaseConfig)
	err = dbc.Connect()
	if err != nil {
		log.Println(err)
		return
	}

	grdbc := grgitrepositorydb.NewClient(dbc)

	grmqc := grscankafka.NewClient(&cfg.GitScannerMQConfig)

	options := &grcmodels.Options{
		Environment:           &cfg,
		FileClient:            fc,
		GitRepositoryDBClient: grdbc,
		GitScannerMQClient:    grmqc,
	}

	// -----------
	// Router
	// -----------
	router := grcrouters.NewWithOptions(options)
	_ = router.Start()
}
