package repository

import (
	"gin-sosmed/entity"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (r *CustomerRepository) Create(customer entity.Customer) error {
	return r.db.Create(&customer).Error
}

func (r *CustomerRepository) Get(id string) (entity.Customer, error) {
	var customer entity.Customer
	err := r.db.First(&customer, "id = ?", id).Error
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (r *CustomerRepository) GetByUserId(id string) (entity.Customer, error) {
	var customer entity.Customer
	err := r.db.First(&customer, "user_id = ?", id).Error
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (r *CustomerRepository) CheckRoom(id string) (bool, error) {
	var room entity.Room
	err := r.db.First(&room, "id = ?", id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *CustomerRepository) CheckStatusRoom(id string) (bool, error) {
	var customer entity.Customer
	err := r.db.First(&customer, "room_id = ? AND check_in = CURRENT_DATE AND check_out = CURRENT_DATE", id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *CustomerRepository) GetAll() (*[]entity.Customer, error) {
	var customer []entity.Customer
	err := r.db.Find(&customer).Error
	return &customer, err
}

func (r *CustomerRepository) Update(p *entity.Customer) (*entity.Customer, error) {
	err := r.db.Save(p).Error
	return p, err
}

func (r *CustomerRepository) Delete(id string) error {
	err := r.db.Delete(&entity.Customer{}, "id = ?", id).Error
	return err
}
