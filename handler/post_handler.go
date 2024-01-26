package handler

import (
	"net/http"
	"os"
	"path/filepath"

	"gin-sosmed/config"
	"gin-sosmed/dto"
	"gin-sosmed/errorhandler"
	"gin-sosmed/helper"
	"gin-sosmed/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type postHandler struct {
	service.PostService
}

func NewPostHandler(p service.PostService) *postHandler {
	return &postHandler{
		PostService: p,
	}
}

func (h *postHandler) Create(c *gin.Context) {
	var post dto.PostRequest

	if err := c.ShouldBind(&post); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.UnprocessableEntityError{Message: err.Error()})
		return
	}

	if post.Photo != nil {
		if err := os.MkdirAll(config.TweetsFolder, 0o755); err != nil {
			errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
				Message: "Can't make tweets folder",
			})
			return
		}

		ext := filepath.Ext(post.Photo.Filename)
		newFileName := uuid.New().String() + ext

		dst := filepath.Join(config.TweetsFolder, newFileName)
		if err := c.SaveUploadedFile(post.Photo, dst); err != nil {
			errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
				Message: "Can't save file to tweets folder",
			})
			return
		}
		post.Photo.Filename = dst
	}

	userID, exist := c.Get("user_id")
	if !exist {
		errorhandler.ErrorHandler(c, &errorhandler.UnauthorizedError{
			Message: "Unauthorize",
		})
		return
	}
	post.AuthorId = userID.(*uuid.UUID)

	if err := h.PostService.Create(&post); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Successfully Created Tweet",
	})
	c.JSON(http.StatusCreated, res)
}
