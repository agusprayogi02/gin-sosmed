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
	var register dto.RegisterRequest

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
