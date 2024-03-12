package main

import (
	"context"

	"gin-sosmed/config"
	"gin-sosmed/router"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	// config
	config.LoadConfig()
	config.LoadDB()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(cors.Default())

	ctx := context.Background()
	l := config.LoadNgrok(ctx)

	// router
	router.InitialRouter(r)

	r.RunListener(l)
	// r.Run(fmt.Sprintf(":%v", config.ENV.PORT))
}
