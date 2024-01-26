package router

import (
	"net/http"

	"gin-sosmed/config"
	"gin-sosmed/middleware"

	"github.com/gin-gonic/gin"
)

func InitialRouter(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ip":      c.ClientIP(),
			"version": gin.Version,
			"url":     "https://above-vulture-monthly.ngrok-free.app",
		})
	})
	// Prevent from calling methods not implemented.
	r.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"status_code": "http.StatusMethodNotAllowed",
			"is_success":  "false",
			"data":        "nil",
			"message":     "Method Not Allowed",
		})
	})

	r.Static(config.TweetUri, config.TweetsFolder)

	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	AuthRouter(api)

	// with auth
	api.Use(middleware.JWTMiddleware())
	PostRouter(api)
}
