package repository

import (
	"gin-sosmed/entity"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) Register(req *entity.User) error {
	err := r.db.Create(&req).Error
	return err
}

func (r *AuthRepository) CreateWithCustomer(userReq *entity.User, customerReq entity.Customer) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&userReq).Error; err != nil {
			return err
		}
		customerReq.UserID = userReq.ID
		if err := tx.Create(&customerReq).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *AuthRepository) EmailExist(email string) bool {
	var user entity.User
	err := r.db.First(&user, "email = ?", email).Error

	return err == nil
}

func (r *AuthRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, "email = ?", email).Error

	return &user, err
}
