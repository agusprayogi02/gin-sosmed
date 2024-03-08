package router

import (
	"gin-sosmed/config"
	"gin-sosmed/handler"
	"gin-sosmed/repository"
	"gin-sosmed/service"

	"github.com/gin-gonic/gin"
)

func NewRoomRouter(r *gin.RouterGroup) {
	repo := repository.NewRoomRepository(config.DB)
	service := service.NewRoomService(repo)
	handler := handler.NewRoomHandler(service)

	api := r.Group("/room")
	api.POST("/", handler.Create)
	api.GET("/:id", handler.Get)
	api.PUT("/:id", handler.Update)
	api.GET("/", handler.GetAll)
}
