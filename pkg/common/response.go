package common

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Code    int    `json:"code"` // HTTP or custom code
	Message string `json:"message,omitempty"`
	Result  any    `json:"result,omitempty"` // 성공시 데이터
	Error   any    `json:"error,omitempty"`  // 실패시 상세 에러 또는 검증정보 등
}

func Success(c *gin.Context, data any, msg ...string) {
	message := ""
	if len(msg) > 0 {
		message = msg[0]
	}
	c.JSON(200, APIResponse{
		Code:    200,
		Message: message,
		Result:  data,
	})
}

func Fail(c *gin.Context, err error) {
	c.Error(err)

	code := http.StatusInternalServerError
	message := "server error"
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		code = apiErr.Code
		message = apiErr.Message
	}

	switch {
	case code >= 500:
		if message == "" {
			message = "server error"
		}
	case code >= 400:
		if message == "" || message == "server error" {
			message = "client error"
		}
	}

	c.AbortWithStatusJSON(code, APIResponse{
		Code:    code,
		Message: message,
	})
}
