package handler

import (
	"go-gin-boilerplate/pkg/common"

	"github.com/gin-gonic/gin"
)

func NewMonitorHandler(r *gin.Engine) {
	group := r.Group("/monitor")

	group.GET("liveness", getLiveness)
}

func getLiveness(c *gin.Context) {
	common.Success(c, true)
}
