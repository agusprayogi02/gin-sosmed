package repository

import (
	"gin-sosmed/dto"
	"gin-sosmed/entity"

	"gorm.io/gorm"
)

type WismaRepository struct {
	db *gorm.DB
}

func NewWismaRepository(db *gorm.DB) *WismaRepository {
	return &WismaRepository{
		db: db,
	}
}

func (r *WismaRepository) Create(wisma entity.Wisma) error {
	return r.db.Create(&wisma).Error
}

func (r *WismaRepository) Get(id string) (entity.Wisma, error) {
	var wisma entity.Wisma
	err := r.db.Preload("User").First(&wisma, "id = ?", id).Error

	return wisma, err
}

func (r *WismaRepository) GetByUser(userId string) (*[]entity.Wisma, error) {
	var wisma []entity.Wisma
	err := r.db.Find(&wisma, "user_id = ?", userId).Error

	return &wisma, err
}

func (r *WismaRepository) GetAll(p *dto.PaginateRequest) (*[]entity.Wisma, error) {
	var wisma []entity.Wisma
	offset := (p.Page - 1) * p.Limit
	err := r.db.Preload("User").Model(&entity.Wisma{}).Limit(p.Limit).Offset(offset).Find(&wisma).Error
	return &wisma, err
}

func (r *WismaRepository) Counter() (int64, error) {
	var count int64
	err := r.db.Model(&entity.Wisma{}).Count(&count).Error
	return count, err
}

func (r *WismaRepository) Update(wisma *entity.Wisma) (*entity.Wisma, error) {
	err := r.db.Updates(wisma).Error
	return wisma, err
}

func (r *WismaRepository) Delete(id string) error {
	var wisma entity.Wisma
	err := r.db.Where("id = ?", id).Delete(&wisma).Error
	if err != nil {
		return err
	}
	return nil
}
