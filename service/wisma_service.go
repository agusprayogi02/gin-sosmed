package service

import (
	"gin-sosmed/dto"
	"gin-sosmed/entity"
	"gin-sosmed/errorhandler"
	"gin-sosmed/repository"

	"github.com/google/uuid"
)

type WismaService interface {
	Create(w dto.WismaRequest) error
	Get(id string) (*dto.WismaResponse, error)
	GetAll(p *dto.PaginateRequest) (*int64, *[]dto.WismaResponse, error)
	Update(id string, req dto.WismaRequest) (*dto.WismaResponse, error)
	Delete(id string) error
}

type wismaService struct {
	repo repository.WismaRepository
}

func NewWismaService(r repository.WismaRepository) *wismaService {
	return &wismaService{
		repo: r,
	}
}

func (s *wismaService) Create(w dto.WismaRequest) error {
	wisma := entity.Wisma{
		ID:      uuid.New(),
		Name:    w.Nama,
		Address: w.Address,
		Code:    w.Code,
		Note:    w.Note,
		UserID:  &w.UserID,
	}

	err := s.repo.Create(wisma)
	if err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return nil
}

func (s *wismaService) Get(id string) (*dto.WismaResponse, error) {
	var wisma *dto.WismaResponse

	data, err := s.repo.Get(id)
	if err != nil {
		return nil, &errorhandler.NotFoundError{
			Message: err.Error(),
		}
	}
	wisma = &dto.WismaResponse{
		ID:      data.ID,
		Nama:    data.Name,
		Address: data.Address,
		Code:    data.Code,
		Note:    data.Note,
		UserID:  *data.UserID,
		User: dto.User{
			ID:    data.User.ID.String(),
			Name:  data.User.Name,
			Email: data.User.Email,
		},
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return wisma, nil
}

func (s *wismaService) GetAll(p *dto.PaginateRequest) (*int64, *[]dto.WismaResponse, error) {
	var data []dto.WismaResponse

	wisma, err := s.repo.GetAll(p)
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}
	count, err := s.repo.Counter()
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}
	for _, w := range *wisma {
		data = append(data, dto.WismaResponse{
			ID:      w.ID,
			Nama:    w.Name,
			Address: w.Address,
			Code:    w.Code,
			Note:    w.Note,
			UserID:  *w.UserID,
			User: dto.User{
				ID:    w.User.ID.String(),
				Name:  w.User.Name,
				Email: w.User.Email,
			},
			CreatedAt: w.CreatedAt,
			UpdatedAt: w.UpdatedAt,
		})
	}

	return &count, &data, nil
}

func (s *wismaService) Update(id string, req dto.WismaRequest) (*dto.WismaResponse, error) {
	var wisma dto.WismaResponse

	data, err := s.repo.Get(id)
	if err != nil {
		return nil, &errorhandler.NotFoundError{
			Message: err.Error(),
		}
	}

	data.Name = req.Nama
	data.Address = req.Address
	data.Code = req.Code
	data.Note = req.Note

	model, err := s.repo.Update(&data)
	if err != nil {
		return nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	wisma = dto.WismaResponse{
		ID:      model.ID,
		Nama:    model.Name,
		Address: model.Address,
		Code:    model.Code,
		Note:    model.Note,
		UserID:  *model.UserID,
		User: dto.User{
			ID:    model.User.ID.String(),
			Name:  model.User.Name,
			Email: model.User.Email,
		},
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
	return &wisma, nil
}

func (s *wismaService) Delete(id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}
	return nil
}
