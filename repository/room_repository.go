package repository

import (
	"gin-sosmed/dto"
	"gin-sosmed/entity"

	"gorm.io/gorm"
)

type RoomRepository interface {
	Create(room entity.Room) error
	Get(id string) (entity.Room, error)
	GetAll(p *dto.PaginateRequest) (*[]entity.Room, error)
	GetByWisma(p *dto.RoomPaginateRequest) (*[]entity.Room, error)
	Counter() (int64, error)
	Update(p *entity.Room) (*entity.Room, error)
	Delete(id string) error
}

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *roomRepository {
	return &roomRepository{
		db: db,
	}
}

func (r *roomRepository) Create(room entity.Room) error {
	return r.db.Create(&room).Error
}

func (r *roomRepository) Get(id string) (entity.Room, error) {
	var room entity.Room
	err := r.db.Preload("Wisma").First(&room, "id = ?", id).Error
	if err != nil {
		return room, err
	}
	return room, nil
}

func (r *roomRepository) GetAll(p *dto.PaginateRequest) (*[]entity.Room, error) {
	var room []entity.Room
	offset := (p.Page - 1) * p.Limit
	err := r.db.Preload("Wisma").Model(&entity.Room{}).Limit(p.Limit).Offset(offset).Find(&room).Error
	return &room, err
}

func (r *roomRepository) GetByWisma(p *dto.RoomPaginateRequest) (*[]entity.Room, error) {
	var room []entity.Room
	offset := (p.Page - 1) * p.Limit
	err := r.db.Preload("Wisma").Model(&entity.Room{}).Limit(p.Limit).Offset(offset).Find(&room, "wisma.id = ?", p.WismaID).Error
	return &room, err
}

func (r *roomRepository) Counter() (int64, error) {
	var count int64
	err := r.db.Model(&entity.Room{}).Count(&count).Error
	return count, err
}

func (r *roomRepository) Update(room *entity.Room) (*entity.Room, error) {
	err := r.db.Updates(room).Error
	return room, err
}

func (r *roomRepository) Delete(id string) error {
	var room entity.Room
	err := r.db.Where("id = ?", id).Delete(&room).Error
	return err
}
