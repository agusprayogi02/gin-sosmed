package errorhandler

import (
	"gin-sosmed/dto"
	"gin-sosmed/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case *NotFoundError:
		statusCode = http.StatusNotFound
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *UnauthorizedError:
		statusCode = http.StatusUnauthorized
	}

	response := helper.Response(dto.ResponseParams{
		StatusCode: statusCode,
		Message:    err.Error(),
	})

	c.JSON(statusCode, response)
}
