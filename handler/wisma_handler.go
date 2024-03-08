package handler

import (
	"net/http"

	"gin-sosmed/config"
	"gin-sosmed/dto"
	"gin-sosmed/errorhandler"
	"gin-sosmed/helper"
	"gin-sosmed/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type wismaHandler struct {
	service service.WismaService
}

func NewWismaHandler(p service.WismaService) *wismaHandler {
	return &wismaHandler{
		service: p,
	}
}

func (h *wismaHandler) Create(c *gin.Context) {
	var req dto.WismaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.UnprocessableEntityError{Message: err.Error()})
		return
	}

	userID, exist := c.Get(config.UserID)
	if exist {
		req.UserID = userID.(uuid.UUID)
	}

	err := h.service.Create(req)
	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Successfully Created Wisma",
	})
	c.JSON(http.StatusCreated, res)
}

func (h *wismaHandler) Get(c *gin.Context) {
	id := c.Param("id")

	data, err := h.service.Get(id)
	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Data:       data,
	})
	c.JSON(http.StatusOK, res)
}

func (h *wismaHandler) GetAll(c *gin.Context) {
	paginate := dto.PaginateRequest{
		Page:  1,  // Default page number
		Limit: 10, // Default limit
	}

	if err := c.ShouldBind(&paginate); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.UnprocessableEntityError{Message: err.Error()})
		return
	}

	total, data, err := h.service.GetAll(&paginate)
	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	res := helper.Response(
		dto.ResponseParams{
			StatusCode: http.StatusOK,
			Data:       data,
			Paginate: &dto.Paginate{
				Page:     paginate.Page,
				PerPage:  paginate.Limit,
				Total:    int(*total),
				NextPage: int(*total) > (paginate.Limit * paginate.Page),
			},
		},
	)
	c.JSON(http.StatusOK, res)
}

func (h *wismaHandler) Update(c *gin.Context) {
	var req dto.WismaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.UnprocessableEntityError{Message: err.Error()})
		return
	}

	id := c.Param("id")
	wisma, err := h.service.Update(id, req)
	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Data:       wisma,
		Message:    "Successfully Updated Wisma",
	})
	c.JSON(http.StatusOK, res)
}

func (h *wismaHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.service.Delete(id)
	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Successfully Deleted Wisma",
	})
	c.JSON(http.StatusOK, res)
}
