package service

import (
	"gin-sosmed/config"
	"gin-sosmed/dto"
	"gin-sosmed/entity"
	"gin-sosmed/errorhandler"
	"gin-sosmed/helper"
	"gin-sosmed/repository"

	"github.com/google/uuid"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
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

	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	user := entity.User{
		ID:       id,
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

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse

	user, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Email not found!"}
	}

	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Wrong password!"}
	}

	token, err := helper.GenerateToken(user, config.ENV.JWT_SIGNING_KEY)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	data = dto.LoginResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return &data, nil
}
