package main

import (
	"context"
	"log"

	"github.com/bersennaidoo/farmstyle/application/rest/handlers"
	"github.com/bersennaidoo/farmstyle/application/rest/mid"
	"github.com/bersennaidoo/farmstyle/application/rest/server"
	"github.com/bersennaidoo/farmstyle/foundation/token"
	"github.com/bersennaidoo/farmstyle/foundation/util"
	"github.com/bersennaidoo/farmstyle/infrastructure/repositories/mongo"
	"github.com/bersennaidoo/farmstyle/physical/config"
	"github.com/bersennaidoo/farmstyle/physical/dbc"
	"github.com/bersennaidoo/farmstyle/physical/logger"
	"github.com/bersennaidoo/farmstyle/physical/otelem"
)

func main() {
	tp, err := otelem.InitTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	log := logger.New()
	filename := config.GetConfigFileName()
	cfg := config.New(filename)
	mclient := dbc.New(cfg)
	pmaker, _ := token.NewPasetoMaker(util.RandomString(32))
	usrepo := mongo.NewUserRepository(mclient, pmaker)
	rvrepo := mongo.NewReviewsRepository(mclient)
	uh := handlers.NewUserHandler(usrepo, log)
	rh := handlers.NewReviewsHandler(rvrepo, log)
	m := mid.New(log, pmaker)
	srv := server.New(uh, rh, cfg, log, m)
	srv.InitRouter()

	log.Info("Starting the application...")
	srv.Start()

	defer func() {
		if err := mclient.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
}
