package router

import (
	"gin-sosmed/config"
	"gin-sosmed/handler"
	"gin-sosmed/repository"
	"gin-sosmed/service"

	"github.com/gin-gonic/gin"
)

func PostRouter(r *gin.RouterGroup) {
	repo := repository.NewPostRepository(config.DB)
	service := service.NewPostService(repo)
	handler := handler.NewPostHandler(service)

	api := r.Group("/tweets")
	api.POST("/", handler.Create)
	api.GET("/:id", handler.Get)
	api.GET("/", handler.GetAll)
}
