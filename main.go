package main

import (
	"fmt"

	"gin-sosmed/config"
	"gin-sosmed/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// config
	config.LoadConfig()
	config.LoadDB()

	r := gin.Default()
	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// router
	router.AuthRouter(api)

	r.Run(fmt.Sprintf(":%v", config.ENV.PORT))
}
