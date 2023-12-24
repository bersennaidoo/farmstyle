package main

import (
	"github.com/bersennaidoo/farmstyle/application/rest/handlers"
	"github.com/bersennaidoo/farmstyle/application/rest/server"
	"github.com/bersennaidoo/farmstyle/infrastructure/repositories/mongo"
	"github.com/bersennaidoo/farmstyle/physical/config"
	"github.com/bersennaidoo/farmstyle/physical/dbc"
	"github.com/bersennaidoo/farmstyle/physical/logger"
)

func main() {

	log := logger.New()
	filename := config.GetConfigFileName()
	cfg := config.New(filename)
	mclient := dbc.New(cfg)
	usrepo := mongo.NewUserRepository(mclient)
	rvrepo := mongo.NewReviewsRepository(mclient)
	uh := handlers.NewUserHandler(usrepo, log)
	rh := handlers.NewReviewsHandler(rvrepo, log)
	srv := server.New(uh, rh, cfg, log)
	srv.InitRouter()

	log.Info("Starting the application...")
	srv.Start()

}
