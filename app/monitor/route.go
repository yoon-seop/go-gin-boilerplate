package monitor

import (
	"github.com/gin-gonic/gin"
	"go-gin-boilerplate/common"
)

func RegisterRoutes(r *gin.Engine) {
	group := r.Group("/monitor")

	group.GET("liveness", getLiveness)
}

func getLiveness(c *gin.Context) {
	common.Success(c, true)
}
