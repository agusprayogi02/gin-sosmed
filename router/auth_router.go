package router

import (
	"gin-sosmed/config"
	"gin-sosmed/handler"
	"gin-sosmed/repository"
	"gin-sosmed/service"

	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup) {
	authRouter := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthService(authRouter)
	authHandler := handler.NewAuthHandler(authService)

	api.POST("/register", authHandler.Register)
}
