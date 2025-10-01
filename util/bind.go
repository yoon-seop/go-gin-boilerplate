package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-boilerplate/common"
	"net/http"
	"strconv"
)

func ParseUint64PathParam(c *gin.Context, param string) (uint64, bool) {
	str := c.Param(param)
	id, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		common.Fail(c, &common.APIError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("`%s` must be a valid positive number", param),
		})

		return 0, false
	}

	return id, true
}
