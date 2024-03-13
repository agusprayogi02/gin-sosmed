package handler

import (
	"net/http"

	"gin-sosmed/dto"
	"gin-sosmed/errorhandler"
	"gin-sosmed/helper"
	"gin-sosmed/service"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

func (h *authHandler) Register(ctx *gin.Context) {
	register := dto.RegisterRequest{
		Role: "admin",
	}

	if err := ctx.ShouldBindJSON(&register); err != nil {
		errorhandler.ErrorHandler(ctx, &errorhandler.UnprocessableEntityError{Message: err.Error()})
		return
	}

	if err := h.service.Register(&register); err != nil {
		errorhandler.ErrorHandler(ctx, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Register Successfully, please login!",
	})
	ctx.JSON(http.StatusCreated, res)
}

func (h *authHandler) Login(ctx *gin.Context) {
	var login dto.LoginRequest

	if err := ctx.ShouldBindJSON(&login); err != nil {
		errorhandler.ErrorHandler(ctx, &errorhandler.UnprocessableEntityError{Message: err.Error()})
		return
	}

	result, err := h.service.Login(&login)
	if err != nil {
		errorhandler.ErrorHandler(ctx, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Login succesfuly",
		Data:       result,
	})
	ctx.JSON(http.StatusOK, res)
}

func (h *authHandler) RegisterCustomer(ctx *gin.Context) {
	var req dto.RegisterCustomerRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		errorhandler.ErrorHandler(ctx, &errorhandler.UnprocessableEntityError{Message: err.Error()})
		return
	}

	err := h.service.RegisterCustomer(&req)
	if err != nil {
		errorhandler.ErrorHandler(ctx, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Register Customer Successfully, please login!",
	})
	ctx.JSON(http.StatusCreated, res)
}
