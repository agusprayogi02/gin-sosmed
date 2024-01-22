package main

import (
	"fmt"
	"gin-sosmed/config"

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
	r.Run(fmt.Sprintf(":%v", config.ENV.PORT))
}
