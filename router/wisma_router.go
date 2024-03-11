package router

import (
	"gin-sosmed/config"
	"gin-sosmed/handler"
	"gin-sosmed/repository"
	"gin-sosmed/service"

	"github.com/gin-gonic/gin"
)

func WismaRouter(r *gin.RouterGroup) {
	repo := repository.NewWismaRepository(config.DB)
	service := service.NewWismaService(*repo)
	handler := handler.NewWismaHandler(*service)

	wr := r.Group("/wisma")
	wr.POST("/", handler.Create)
	wr.GET("/:id", handler.Get)
	wr.PUT("/:id", handler.Update)
	wr.GET("/", handler.GetAll)
	wr.DELETE("/:id", handler.Delete)
}
