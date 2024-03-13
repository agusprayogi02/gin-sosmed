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

type AuthService struct {
	repository repository.AuthRepository
	wismaRepo  repository.WismaRepository
}

func NewAuthService(r repository.AuthRepository, w repository.WismaRepository) *AuthService {
	return &AuthService{
		repository: r,
		wismaRepo:  w,
	}
}

func (s *AuthService) RegisterCheck(req *dto.RegisterRequest) (*entity.User, error) {
	if emailExist := s.repository.EmailExist(req.Email); emailExist {
		return nil, &errorhandler.UnprocessableEntityError{
			Message: "Email already exist",
		}
	}

	if req.Password != req.PasswordConfirm {
		return nil, &errorhandler.UnprocessableEntityError{
			Message: "Password not same",
		}
	}

	pass, err := helper.HashPassword(req.Password)
	if err != nil {
		return nil, &errorhandler.InternalServerError{
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
		id = uuid.New()
	}

	var role entity.RoleType
	if req.Role == entity.ADMIN.String() {
		role = entity.ADMIN
	} else {
		role = entity.USER
	}

	user := entity.User{
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		Password: pass,
		Gender:   gender,
		Role:     role,
	}
	return &user, nil
}

func (s *AuthService) Register(req *dto.RegisterRequest) error {
	user, err := s.RegisterCheck(req)
	if err != nil {
		return err
	}
	if err := s.repository.Register(user); err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return nil
}

func (s *AuthService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
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

	var wismasRes []dto.WismaResponse
	wismas, err := s.wismaRepo.GetByUser(user.ID.String())
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}
	for _, v := range *wismas {
		wismasRes = append(wismasRes, dto.WismaResponse{
			ID:      v.ID,
			Nama:    v.Name,
			Address: v.Address,
			Code:    v.Code,
			UserID:  *v.UserID,
			User: &dto.User{
				ID:    user.ID.String(),
				Name:  user.Name,
				Email: user.Email,
			},
			Note: v.Note,
		})
	}

	data = dto.LoginResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
		Role:  user.Role.String(),
		Wisma: &wismasRes,
	}

	return &data, nil
}

func (s *AuthService) RegisterCustomer(req *dto.RegisterCustomerRequest) error {
	userReq := dto.RegisterRequest{
		Name:            req.Name,
		Email:           req.Email,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
		Gender:          req.Gender,
		Role:            "user",
	}

	user, err := s.RegisterCheck(&userReq)
	if err != nil {
		return err
	}

	customerReq := entity.Customer{
		ID:      uuid.New(),
		Nik:     req.NIK,
		Name:    req.Name,
		Address: &req.Address,
		Phone:   req.Phone,
		UserID:  user.ID,
	}
	if err := s.repository.CreateWithCustomer(user, customerReq); err != nil {
		return &errorhandler.NotFoundError{Message: err.Error()}
	}

	return nil
}
