package user

import (
	"github.com/gin-gonic/gin"
	"go-gin-boilerplate/common"
	"go-gin-boilerplate/util"
	"net/http"
)

func RegisterRoutes(r *gin.Engine) {
	group := r.Group("/user")

	group.POST("", postUser)
	group.GET(":id", getUser)
}

func postUser(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Fail(c, &common.APIError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	user, err := CreateUser(req)
	if err != nil {
		common.Fail(c, err)
		return
	}

	common.Success(c, ToV1Response(user))
}

func getUser(c *gin.Context) {
	id, ok := util.ParseUint64PathParam(c, "id")
	if !ok {
		return
	}

	user, err := GetUserByID(id)
	if err != nil {
		common.Fail(c, err)
		return
	}

	common.Success(c, ToV1Response(user))
}
