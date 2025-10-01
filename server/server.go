package server

import (
	"github.com/gin-gonic/gin"
	"go-gin-boilerplate/app/monitor"
	"go-gin-boilerplate/app/user"
	"go-gin-boilerplate/config"
	"go-gin-boilerplate/logger"
)

func NewEngine(_ *config.AppConfig) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(logger.ZerologMiddleware())

	monitor.RegisterRoutes(router)
	user.RegisterRoutes(router)

	return router
}
