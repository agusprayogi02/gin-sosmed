package handler

import (
	"fmt"
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
	service service.PostService
}

func NewPostHandler(p service.PostService) *postHandler {
	return &postHandler{
		service: p,
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
		post.Photo.Filename = config.TweetUri + newFileName
	}

	userID, exist := c.Get(config.UserID)
	if exist {
		post.AuthorId = userID.(uuid.UUID)
	}

	if err := h.service.Create(&post); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Successfully Created Tweet",
	})
	c.JSON(http.StatusCreated, res)
}

func (h *postHandler) Get(c *gin.Context) {
	id := c.Param("id")
	post, err := h.service.Get(id)
	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}
	tempPhoto := fmt.Sprintf("http://%v/%v", c.Request.Host, *post.Photo)
	post.Photo = &tempPhoto
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Data:       post,
	})

	c.JSON(http.StatusOK, res)
}

func (h *postHandler) GetAll(c *gin.Context) {
	var paginate *dto.PaginateRequest
	if err := c.ShouldBindQuery(&paginate); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}
	total, data, err := h.service.GetAll(paginate)
	if err != nil {
		errorhandler.ErrorHandler(c, err)
	}

	res := helper.Response(
		dto.ResponseParams{
			StatusCode: http.StatusOK,
			Message:    "Berhasil",
			Paginate: &dto.Paginate{
				Page:    paginate.Page,
				PerPage: paginate.Limit,
				Total:   int(*total),
			},
			Data: data,
		},
	)
	c.JSON(http.StatusOK, res)
}
