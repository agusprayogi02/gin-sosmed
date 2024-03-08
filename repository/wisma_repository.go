package repository

import (
	"gin-sosmed/dto"
	"gin-sosmed/entity"

	"gorm.io/gorm"
)

type WismaRepository interface {
	Create(wisma entity.Wisma) error
	Get(id string) (entity.Wisma, error)
	GetAll(p *dto.PaginateRequest) (*[]entity.Wisma, error)
	Counter() (int64, error)
	Update(p *entity.Wisma) (*entity.Wisma, error)
	Delete(id string) error
}

type wismaRepository struct {
	db *gorm.DB
}

func NewWismaRepository(db *gorm.DB) *wismaRepository {
	return &wismaRepository{
		db: db,
	}
}

func (r *wismaRepository) Create(wisma entity.Wisma) error {
	return r.db.Create(&wisma).Error
}

func (r *wismaRepository) Get(id string) (entity.Wisma, error) {
	var wisma entity.Wisma
	err := r.db.Preload("User").First(&wisma, "id = ?", id).Error

	return wisma, err
}

func (r *wismaRepository) GetAll(p *dto.PaginateRequest) (*[]entity.Wisma, error) {
	var wisma []entity.Wisma
	offset := (p.Page - 1) * p.Limit
	err := r.db.Preload("User").Model(&entity.Wisma{}).Limit(p.Limit).Offset(offset).Find(&wisma).Error
	return &wisma, err
}

func (r *wismaRepository) Counter() (int64, error) {
	var count int64
	err := r.db.Model(&entity.Wisma{}).Count(&count).Error
	return count, err
}

func (r *wismaRepository) Update(wisma *entity.Wisma) (*entity.Wisma, error) {
	err := r.db.Updates(wisma).Error
	return wisma, err
}

func (r *wismaRepository) Delete(id string) error {
	var wisma entity.Wisma
	err := r.db.Where("id = ?", id).Delete(&wisma).Error
	if err != nil {
		return err
	}
	return nil
}
