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

func (s *CustomerService) Scan(req *dto.CustomerScan) error {
	roomId, err := uuid.Parse(req.RoomID)
	if err != nil {
		return &errorhandler.UnprocessableEntityError{
			Message: "Invalid Room ID",
		}
	}
	customer, err := s.repo.GetByUserId(req.UserID.String())
	if err != nil {
		return &errorhandler.NotFoundError{
			Message: "Customer Not Found",
		}
	}

	room, _ := s.repo.CheckRoom(req.RoomID)
	if !room {
		return &errorhandler.NotFoundError{
			Message: "Room Not Found",
		}
	}
	status, _ := s.repo.CheckStatusRoom(req.RoomID)
	if status {
		return &errorhandler.UnprocessableEntityError{
			Message: "Room already booked",
		}
	}

	now := time.Now().Format("2006-01-02")
	checkOut, err := time.Parse("2006-01-02", req.CheckOut)
	if err != nil {
		return &errorhandler.UnprocessableEntityError{
			Message: "Invalid Check Out Date",
		}
	}
	checkIn, err := time.Parse("2006-01-02", now)
	if err != nil {
		return &errorhandler.UnprocessableEntityError{
			Message: "Invalid Check Out Date",
		}
	}
	if checkIn.After(checkOut) {
		return &errorhandler.UnprocessableEntityError{
			Message: "Check Out Date must be greater than Check In Date",
		}
	}

	customer.RoomID = &roomId
	customer.CheckIn = &checkIn
	customer.CheckOut = &checkOut

	_, err = s.repo.Update(&customer)
	if err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}
	return nil
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

	return &dto.CustomerResponse{
		ID:       customer.ID,
		Nik:      customer.Nik,
		Name:     customer.Name,
		Address:  customer.Address,
		Phone:    customer.Phone,
		RoomID:   *customer.RoomID,
		UserID:   customer.UserID,
		CheckIn:  customer.CheckIn,
		CheckOut: customer.CheckOut,
	}, nil
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
		response = append(response, dto.CustomerResponse{
			ID:       customer.ID,
			Nik:      customer.Nik,
			Name:     customer.Name,
			Address:  customer.Address,
			Phone:    customer.Phone,
			RoomID:   *customer.RoomID,
			UserID:   customer.UserID,
			CheckIn:  customer.CheckIn,
			CheckOut: customer.CheckOut,
		})
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

	return &dto.CustomerResponse{
		ID:       updatedCustomer.ID,
		Nik:      updatedCustomer.Nik,
		Name:     updatedCustomer.Name,
		Address:  updatedCustomer.Address,
		Phone:    updatedCustomer.Phone,
		RoomID:   *updatedCustomer.RoomID,
		UserID:   updatedCustomer.UserID,
		CheckIn:  updatedCustomer.CheckIn,
		CheckOut: updatedCustomer.CheckOut,
	}, nil
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
