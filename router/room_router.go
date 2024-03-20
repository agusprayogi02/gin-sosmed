package router

import (
	"gin-sosmed/config"
	"gin-sosmed/handler"
	"gin-sosmed/repository"
	"gin-sosmed/service"

	"github.com/gin-gonic/gin"
)

func RoomRouter(r *gin.RouterGroup) {
	repo := repository.NewRoomRepository(config.DB)
	service := service.NewRoomService(*repo)
	handler := handler.NewRoomHandler(*service)

	api := r.Group("/room")
	api.POST("", handler.Create)
	api.GET("/:id", handler.Get)
	api.DELETE("/:id", handler.Delete)
	api.PUT("/:id", handler.Update)
	api.GET("", handler.GetAll)
	api.GET("/wisma", handler.GetByWisma)
	api.GET("/user", handler.GetByUser)
	api.GET("/user-raw", handler.GetByUserRaw)
}
