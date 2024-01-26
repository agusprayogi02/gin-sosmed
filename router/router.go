package router

import (
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
