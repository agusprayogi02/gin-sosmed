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

type CustomerHandler struct {
	service *service.CustomerService
}

func NewCustomerHandler(s *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		service: s,
	}
}

func (h *CustomerHandler) ScanQr(c *gin.Context) {
	var req dto.CustomerScan

	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.UnprocessableEntityError{Message: err.Error()})
		return
	}

	userID, exist := c.Get(config.UserID)
	if !exist {
		errorhandler.ErrorHandler(c, &errorhandler.UnauthorizedError{Message: "Unauthorized"})
		return
	}

	req.UserID = userID.(uuid.UUID)
	data, err := h.service.Scan(&req)
	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Successfully Booked Room",
		Data:       data,
	})
	c.JSON(http.StatusCreated, res)
}

func (h *CustomerHandler) Create(c *gin.Context) {
	var req dto.CustomerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.UnprocessableEntityError{Message: err.Error()})
		return
	}

	userID, exist := c.Get(config.UserID)
	if !exist {
		errorhandler.ErrorHandler(c, &errorhandler.UnauthorizedError{Message: "Unauthorized"})
		return
	}
	req.UserID = userID.(uuid.UUID)

	err := h.service.Create(&req)
	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Successfully Created Customer",
	})
	c.JSON(http.StatusCreated, res)
}

func (h *CustomerHandler) Get(c *gin.Context) {
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

func (h *CustomerHandler) GetAll(c *gin.Context) {
	data, err := h.service.GetAll()
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

func (h *CustomerHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req dto.CustomerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.UnprocessableEntityError{Message: err.Error()})
		return
	}

	userID, exist := c.Get(config.UserID)
	if !exist {
		errorhandler.ErrorHandler(c, &errorhandler.UnauthorizedError{Message: "Unauthorized"})
		return
	}
	req.UserID = userID.(uuid.UUID)

	data, err := h.service.Update(&req, id)
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

func (h *CustomerHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.service.Delete(id)
	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Successfully Deleted Customer",
	})
	c.JSON(http.StatusOK, res)
}
