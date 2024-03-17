package router

import (
	"gin-sosmed/config"
	"gin-sosmed/handler"
	"gin-sosmed/repository"
	"gin-sosmed/service"

	"github.com/gin-gonic/gin"
)

func CustomerRouter(r *gin.RouterGroup) {
	repo := repository.NewCustomerRepository(config.DB)
	service := service.NewCustomerService(repo)
	handler := handler.NewCustomerHandler(service)

	api := r.Group("/customer")
	api.POST("", handler.Create)
	api.GET("/user", handler.GetByUser)
	api.POST("/scan", handler.ScanQr)
	api.GET("/:id", handler.Get)
	api.PUT("/:id", handler.Update)
	api.GET("", handler.GetAll)
	api.DELETE("/:id", handler.Delete)
}
