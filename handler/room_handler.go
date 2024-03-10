package handler

import (
	"net/http"

	"gin-sosmed/dto"
	"gin-sosmed/errorhandler"
	"gin-sosmed/helper"
	"gin-sosmed/service"

	"github.com/gin-gonic/gin"
)

type roomHandler struct {
	service service.RoomService
}

func NewRoomHandler(p service.RoomService) *roomHandler {
	return &roomHandler{
		service: p,
	}
}

func (h *roomHandler) Create(c *gin.Context) {
	req := dto.RoomRequest{
		Capacity: 1,
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.UnprocessableEntityError{Message: err.Error()})
		return
	}

	if req.Capacity <= 0 {
		errorhandler.ErrorHandler(c, &errorhandler.UnprocessableEntityError{Message: "Kapasitar kamar tidak boleh kurang dari 1"})
		return
	}

	err := h.service.Create(&req)
	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Successfully Created Room",
	})
	c.JSON(http.StatusCreated, res)
}

func (h *roomHandler) Get(c *gin.Context) {
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

func (h *roomHandler) GetAll(c *gin.Context) {
	paginate := dto.PaginateRequest{
		Page:  1,  // Default page number
		Limit: 10, // Default limit
	}

	if err := c.ShouldBind(&paginate); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.UnprocessableEntityError{
			Message: err.Error(),
		})
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

func (h *roomHandler) GetByWisma(c *gin.Context) {
	paginate := dto.RoomPaginateRequest{
		Page:  1,  // Default page number
		Limit: 10, // Default limit
	}

	if err := c.ShouldBind(&paginate); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.UnprocessableEntityError{
			Message: err.Error(),
		})
		return
	}
	total, data, err := h.service.GetByWisma(&paginate)
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

func (h *roomHandler) Update(c *gin.Context) {
	var req dto.RoomEditRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.UnprocessableEntityError{
			Message: err.Error(),
		})
		return
	}

	id := c.Param("id")
	data, err := h.service.Update(id, req)
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

func (h *roomHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.service.Delete(id)
	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Successfully Deleted Room",
	})
	c.JSON(http.StatusOK, res)
}
