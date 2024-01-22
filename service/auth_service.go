package service

import (
	"gin-sosmed/dto"
	"gin-sosmed/entity"
	"gin-sosmed/errorhandler"
	"gin-sosmed/helper"
	"gin-sosmed/repository"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) error {
	if emailExist := s.repository.EmailExist(req.Email); emailExist {
		return &errorhandler.UnprocessableEntityError{
			Message: "Email already exist",
		}
	}

	if req.Password != req.PasswordConfirm {
		return &errorhandler.UnprocessableEntityError{
			Message: "Password not same",
		}
	}

	pass, err := helper.HashPassword(req.Password)
	if err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	var gender entity.GenderType
	if req.Gender == entity.PRIA.String() {
		gender = entity.PRIA
	} else {
		gender = entity.WANITA
	}

	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: pass,
		Gender:   gender,
	}

	if err := s.repository.Register(&user); err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return nil
}
