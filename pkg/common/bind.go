package common

import (
	"fmt"

	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ParseUint64PathParam(c *gin.Context, param string) (uint64, bool) {
	str := c.Param(param)
	id, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		Fail(c, &APIError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("`%s` must be a valid positive number", param),
		})

		return 0, false
	}

	return id, true
}

func BindJSONBody[T any](c *gin.Context, req *T) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorMessage := make([]string, 0)

			for _, fieldError := range validationErrors {
				fullFieldName := fieldError.Field()
				parts := strings.Split(fullFieldName, ".")
				fieldName := parts[len(parts)-1]

				msg := fmt.Sprintf("Field `%s` is invalid (Rule: %s)",
					fieldName,
					fieldError.Tag(),
				)
				errorMessage = append(errorMessage, msg)
			}

			Fail(c, &APIError{
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("Request body validation failed. %s", errorMessage),
			})

			return false
		}

		Fail(c, &APIError{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body format or data type mismatch.",
		})

		return false
	}

	return true
}
