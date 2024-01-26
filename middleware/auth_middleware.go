package middleware

import (
	"strings"

	"gin-sosmed/config"
	"gin-sosmed/errorhandler"
	"gin-sosmed/helper"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || tokenString == "*" {
			errorhandler.ErrorHandler(c, &errorhandler.UnauthorizedError{
				Message: "Unauthorize",
			})
			c.Abort()
			return
		}
		tokenArr := strings.Split(tokenString, " ")

		userID, err := helper.VerifyToken(tokenArr[len(tokenArr)-1], config.ENV.JWT_SIGNING_KEY)
		if err != nil {
			errorhandler.ErrorHandler(c, &errorhandler.UnauthorizedError{
				Message: err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}
