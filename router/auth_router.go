package router

import (
	"gin-sosmed/config"
	"gin-sosmed/handler"
	"gin-sosmed/repository"
	"gin-sosmed/service"

	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup) {
	authRepo := repository.NewAuthRepository(config.DB)
	wismaRepo := repository.NewWismaRepository(config.DB)
	authService := service.NewAuthService(*authRepo, *wismaRepo)
	authHandler := handler.NewAuthHandler(*authService)

	api.POST("/register", authHandler.Register)
	api.POST("/register-customer", authHandler.RegisterCustomer)
	api.POST("/login", authHandler.Login)
}
