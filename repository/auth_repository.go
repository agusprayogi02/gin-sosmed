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
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Create(&userReq).Error; err != nil {
		tx.Rollback()
		return err
	}
	customerReq.UserID = userReq.ID
	if err := tx.Create(&customerReq).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
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
