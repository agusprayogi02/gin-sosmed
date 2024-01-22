package errorhandler

import (
	"net/http"

	"gin-sosmed/dto"
	"gin-sosmed/helper"

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
	case *UnprocessableEntityError:
		statusCode = http.StatusUnprocessableEntity
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: statusCode,
		Message:    err.Error(),
	})
	c.JSON(statusCode, res)
}
