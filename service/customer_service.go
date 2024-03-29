package service

import (
	"time"

	"gin-sosmed/dto"
	"gin-sosmed/entity"
	"gin-sosmed/errorhandler"
	"gin-sosmed/repository"

	"github.com/google/uuid"
)

type CustomerService struct {
	repo *repository.CustomerRepository
}

func NewCustomerService(repo *repository.CustomerRepository) *CustomerService {
	return &CustomerService{
		repo: repo,
	}
}

func (s *CustomerService) Scan(req *dto.CustomerScan) (*dto.CustomerResponse, error) {
	roomId, err := uuid.Parse(req.RoomID)
	if err != nil {
		return nil, &errorhandler.UnprocessableEntityError{
			Message: "Invalid Room ID",
		}
	}

	now := time.Now().Format("2006-01-02")
	checkOut, err := time.Parse("2006-01-02", req.CheckOut)
	if err != nil {
		return nil, &errorhandler.UnprocessableEntityError{
			Message: "Invalid Check Out Date",
		}
	}
	checkIn, err := time.Parse("2006-01-02", now)
	if err != nil {
		return nil, &errorhandler.UnprocessableEntityError{
			Message: "Invalid Check Out Date",
		}
	}
	if checkIn.After(checkOut) {
		return nil, &errorhandler.UnprocessableEntityError{
			Message: "Check Out Date must be greater than Check In Date",
		}
	}

	customer, err := s.repo.GetByUserId(req.UserID.String())
	if err != nil {
		return nil, &errorhandler.NotFoundError{
			Message: "Customer Not Found",
		}
	}

	if customer.RoomID != nil {
		if customer.CheckOut.After(checkIn) {
			return nil, &errorhandler.UnprocessableEntityError{
				Message: "You have booked a room",
			}
		}
	}

	room, _ := s.repo.CheckRoom(req.RoomID)
	if !room {
		return nil, &errorhandler.NotFoundError{
			Message: "Room Not Found",
		}
	}
	status, _ := s.repo.CheckStatusRoom(req.RoomID)
	if status {
		return nil, &errorhandler.UnprocessableEntityError{
			Message: "Room already booked",
		}
	}

	customer.RoomID = &roomId
	customer.CheckIn = &checkIn
	customer.CheckOut = &checkOut

	_, err = s.repo.Update(&customer)
	if err != nil {
		return nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}
	return s.Get(customer.ID.String())
}

func (s *CustomerService) Create(req *dto.CustomerRequest) error {
	roomId, err := uuid.Parse(req.RoomID)
	if err != nil {
		return &errorhandler.UnprocessableEntityError{
			Message: "Invalid Room ID",
		}
	}
	consumer := entity.Customer{
		ID:       uuid.New(),
		Nik:      req.Nik,
		Name:     req.Name,
		Address:  req.Address,
		Phone:    req.Phone,
		RoomID:   &roomId,
		UserID:   req.UserID,
		CheckIn:  &req.CheckIn,
		CheckOut: &req.CheckOut,
	}
	return s.repo.Create(consumer)
}

func (s *CustomerService) Get(id string) (*dto.CustomerResponse, error) {
	customer, err := s.repo.Get(id)
	if err != nil {
		return nil, &errorhandler.NotFoundError{
			Message: "Customer Not Found",
		}
	}
	return s.HandleGet(customer), nil
}

func (s *CustomerService) GetByUserId(id string) (*dto.CustomerResponse, error) {
	customer, err := s.repo.GetByUserId(id)
	if err != nil {
		return nil, &errorhandler.NotFoundError{
			Message: "Customer Not Found",
		}
	}
	if customer.CheckOut != nil && time.Now().After(*customer.CheckOut) {
		customer.Room = nil
		customer.RoomID = nil
		customer.CheckIn = nil
		customer.CheckOut = nil

		_, err = s.repo.Update(&customer)
		if err != nil {
			return nil, &errorhandler.InternalServerError{
				Message: err.Error(),
			}
		}
		customer, err = s.repo.GetByUserId(id)
		if err != nil {
			return nil, &errorhandler.NotFoundError{
				Message: "Customer Not Found",
			}
		}
	}

	return s.HandleGet(customer), nil
}

func (s *CustomerService) HandleGet(customer entity.Customer) *dto.CustomerResponse {
	response := &dto.CustomerResponse{
		ID:      customer.ID,
		Nik:     customer.Nik,
		Name:    customer.Name,
		Address: customer.Address,
		Phone:   customer.Phone,
		RoomID:  customer.RoomID,
		Room:    nil,
		User: dto.User{
			ID:    customer.User.ID.String(),
			Email: customer.User.Email,
			Name:  customer.User.Name,
		},
		UserID:   customer.UserID,
		CheckIn:  customer.CheckIn,
		CheckOut: customer.CheckOut,
	}
	if customer.RoomID != nil {
		response.Room = &dto.RoomResponse{
			ID:      customer.Room.ID,
			Name:    customer.Room.Name,
			WismaID: customer.Room.WismaID,
			Wisma: &dto.WismaResponse{
				ID:        customer.Room.Wisma.ID,
				Name:      customer.Room.Wisma.Name,
				Address:   customer.Room.Wisma.Address,
				Code:      customer.Room.Wisma.Code,
				Note:      customer.Room.Wisma.Note,
				User:      nil,
				CreatedAt: customer.Room.Wisma.CreatedAt,
				UpdatedAt: customer.Room.Wisma.UpdatedAt,
			},
			Capacity:  customer.Room.Capacity,
			Note:      customer.Room.Note,
			CreatedAt: customer.Room.CreatedAt,
			UpdatedAt: customer.Room.UpdatedAt,
		}
	}
	return response
}

func (s *CustomerService) GetAll() (*[]dto.CustomerResponse, error) {
	customers, err := s.repo.GetAll()
	if err != nil {
		return nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	var response []dto.CustomerResponse
	for _, customer := range *customers {
		response = append(response, *s.HandleGet(customer))
	}

	return &response, nil
}

func (s *CustomerService) Update(req *dto.CustomerRequest, id string) (*dto.CustomerResponse, error) {
	roomId, err := uuid.Parse(req.RoomID)
	if err != nil {
		return nil, &errorhandler.UnprocessableEntityError{
			Message: "Invalid Room ID",
		}
	}

	data, err := s.repo.Get(id)
	if err != nil {
		return nil, &errorhandler.NotFoundError{
			Message: "Customer Not Found",
		}
	}

	data.Nik = req.Nik
	data.Name = req.Name
	data.Address = req.Address
	data.Phone = req.Phone
	data.RoomID = &roomId
	data.CheckIn = &req.CheckIn
	data.CheckOut = &req.CheckOut

	updatedCustomer, err := s.repo.Update(&data)
	if err != nil {
		return nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return s.Get(updatedCustomer.ID.String())
}

func (s *CustomerService) Delete(id string) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return &errorhandler.NotFoundError{
			Message: "Customer Not Found",
		}
	}

	return s.repo.Delete(id)
}
