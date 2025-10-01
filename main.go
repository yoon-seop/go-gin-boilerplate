package main

import (
	"flag"
	"go-gin-boilerplate/config"
	"go-gin-boilerplate/internal/handler"
	"go-gin-boilerplate/pkg/logger"
	"log"

	"github.com/gin-gonic/gin"
	zerolog "github.com/rs/zerolog/log"
)

func main() {
	configName := flag.String("config", "development", "config file name (without .yaml)")
	flag.Parse()

	conf, err := config.LoadConfig(*configName)
	if err != nil {
		log.Fatalf("Could not load config: %v\n", err)
	}

	if err := config.DatabaseInitialize(conf); err != nil {
		log.Fatalf("Could not initialize DB: %v\n", err)
	}

	logger.Initialize(conf.Env())

	router := gin.New()

	if conf.Env() == config.Development {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Use(gin.Recovery())
	router.Use(logger.ZerologMiddleware())

	handler.NewMonitorHandler(router)
	handler.NewUserHandler(router)
	handler.NewPostHandler(router)

	if err = router.Run(conf.Port()); err != nil {
		zerolog.Fatal().Err(err).Msg("failed to start server.")
	}

	zerolog.Info().Msg("server on.")
}
