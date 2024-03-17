package service

import (
	"gin-sosmed/dto"
	"gin-sosmed/entity"
	"gin-sosmed/errorhandler"
	"gin-sosmed/repository"

	"github.com/google/uuid"
)

type RoomService struct {
	repo repository.RoomRepository
}

func NewRoomService(r repository.RoomRepository) *RoomService {
	return &RoomService{
		repo: r,
	}
}

func (r *RoomService) Create(req *dto.RoomRequest) error {
	id, err := uuid.Parse(req.WismaID)
	if err != nil {
		return &errorhandler.BadRequestError{
			Message: err.Error(),
		}
	}
	room := entity.Room{
		ID:       uuid.New(),
		Name:     req.Name,
		WismaID:  &id,
		Capacity: req.Capacity,
		Note:     req.Note,
	}

	if err := r.repo.Create(room); err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return nil
}

func (r *RoomService) Get(id string) (*dto.RoomResponse, error) {
	var room *dto.RoomResponse

	data, err := r.repo.Get(id)
	if err != nil {
		return nil, &errorhandler.NotFoundError{
			Message: err.Error(),
		}
	}
	room = &dto.RoomResponse{
		ID:       data.ID,
		Name:     data.Name,
		WismaID:  data.WismaID,
		Capacity: data.Capacity,
		Wisma: &dto.WismaResponse{
			ID:        data.Wisma.ID,
			Name:      data.Wisma.Name,
			Address:   data.Wisma.Address,
			Code:      data.Wisma.Code,
			Note:      data.Wisma.Note,
			UserID:    *data.Wisma.UserID,
			CreatedAt: data.Wisma.CreatedAt,
			UpdatedAt: data.Wisma.UpdatedAt,
		},
		Note: data.Note,
	}

	return room, nil
}

func (r *RoomService) GetAll(p *dto.PaginateRequest) (*int64, *[]dto.RoomResponse, error) {
	var rooms []dto.RoomResponse

	count, err := r.repo.Counter()
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	data, err := r.repo.GetAll(p)
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	for _, v := range *data {
		rooms = append(rooms, dto.RoomResponse{
			ID:       v.ID,
			Name:     v.Name,
			WismaID:  v.WismaID,
			Capacity: v.Capacity,
			Wisma: &dto.WismaResponse{
				ID:        v.Wisma.ID,
				Name:      v.Wisma.Name,
				Address:   v.Wisma.Address,
				Code:      v.Wisma.Code,
				Note:      v.Wisma.Note,
				UserID:    *v.Wisma.UserID,
				CreatedAt: v.Wisma.CreatedAt,
				UpdatedAt: v.Wisma.UpdatedAt,
			},
			Note:      v.Note,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return &count, &rooms, nil
}

func (r *RoomService) GetByWisma(p *dto.RoomPaginateRequest) (*int64, *[]dto.RoomResponse, error) {
	var rooms []dto.RoomResponse

	count, err := r.repo.Counter()
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	data, err := r.repo.GetByWisma(p)
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	for _, v := range *data {
		rooms = append(rooms, dto.RoomResponse{
			ID:       v.ID,
			Name:     v.Name,
			WismaID:  v.WismaID,
			Capacity: v.Capacity,
			Wisma: &dto.WismaResponse{
				ID:        v.Wisma.ID,
				Name:      v.Wisma.Name,
				Address:   v.Wisma.Address,
				Code:      v.Wisma.Code,
				Note:      v.Wisma.Note,
				UserID:    *v.Wisma.UserID,
				CreatedAt: v.Wisma.CreatedAt,
				UpdatedAt: v.Wisma.UpdatedAt,
			},
			Note:      v.Note,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return &count, &rooms, nil
}

func (r *RoomService) GetByUserRaw(p *dto.UserRoomPaginateRequest) (*int64, *[]dto.RoomResponse, error) {
	count, err := r.repo.Counter()
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	rooms, err := r.repo.GetByUserRaw(p)
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}
	return &count, rooms, nil
}

func (r *RoomService) GetByUser(p *dto.UserRoomPaginateRequest) (*int64, *[]dto.RoomResponse, error) {
	var rooms []dto.RoomResponse

	count, err := r.repo.Counter()
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	data, err := r.repo.GetByUser(p)
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	for _, v := range *data {
		rooms = append(rooms, dto.RoomResponse{
			ID:       v.ID,
			Name:     v.Name,
			WismaID:  v.WismaID,
			Capacity: v.Capacity,
			Wisma: &dto.WismaResponse{
				ID:      v.Wisma.ID,
				Name:    v.Wisma.Name,
				Address: v.Wisma.Address,
				Code:    v.Wisma.Code,
				UserID:  *v.Wisma.UserID,
				User: &dto.User{
					ID:    v.Wisma.User.ID.String(),
					Name:  v.Wisma.User.Name,
					Email: v.Wisma.User.Email,
				},
				Note:      v.Wisma.Note,
				CreatedAt: v.Wisma.CreatedAt,
				UpdatedAt: v.Wisma.UpdatedAt,
			},
			Note:      v.Note,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return &count, &rooms, nil
}

func (r *RoomService) Update(id string, req dto.RoomEditRequest) (*dto.RoomResponse, error) {
	data, err := r.repo.Get(id)
	if err != nil {
		return nil, &errorhandler.NotFoundError{
			Message: err.Error(),
		}
	}

	data.Capacity = req.Capacity
	data.Name = req.Name
	data.Note = req.Note

	if err := r.repo.Update(&data); err != nil {
		return nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return r.Get(id)
}

func (r *RoomService) Delete(id string) error {
	err := r.repo.Delete(id)
	if err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return nil
}
