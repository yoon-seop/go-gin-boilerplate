package handler

import (
	"go-gin-boilerplate/internal/app/user"
	"go-gin-boilerplate/pkg/common"

	"github.com/gin-gonic/gin"
)

func NewUserHandler(r *gin.Engine) {
	group := r.Group("/user")

	group.POST("", postUser)
	group.GET(":id", getUser)
}

func postUser(c *gin.Context) {
	var req user.CreateRequest
	if ok := common.BindJSONBody(c, &req); !ok {
		return
	}

	data, err := user.CreateUser(c, req)
	if err != nil {
		common.Fail(c, err)
		return
	}

	response := user.ToV1(data)

	common.Success(c, response)
}

func getUser(c *gin.Context) {
	id, ok := common.ParseUint64PathParam(c, "id")
	if !ok {
		return
	}

	data, err := user.GetUserByID(c, id)
	if err != nil {
		common.Fail(c, err)
		return
	}

	response := user.ToV1(data)

	common.Success(c, response)
}
