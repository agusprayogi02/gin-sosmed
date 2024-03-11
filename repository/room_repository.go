package repository

import (
	"gin-sosmed/dto"
	"gin-sosmed/entity"

	"gorm.io/gorm"
)

type RoomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{
		db: db,
	}
}

func (r RoomRepository) Create(room entity.Room) error {
	return r.db.Create(&room).Error
}

func (r RoomRepository) Get(id string) (entity.Room, error) {
	var room entity.Room
	err := r.db.Preload("Wisma").First(&room, "id = ?", id).Error
	if err != nil {
		return room, err
	}
	return room, nil
}

func (r RoomRepository) GetAll(p *dto.PaginateRequest) (*[]entity.Room, error) {
	var room []entity.Room
	offset := (p.Page - 1) * p.Limit
	err := r.db.Preload("Wisma").Model(&entity.Room{}).Limit(p.Limit).Offset(offset).Find(&room).Error
	return &room, err
}

func (r RoomRepository) GetByWisma(p *dto.RoomPaginateRequest) (*[]entity.Room, error) {
	var room []entity.Room
	offset := (p.Page - 1) * p.Limit
	err := r.db.Preload("Wisma").Model(&entity.Room{}).Limit(p.Limit).Offset(offset).Find(&room, "wisma_id = ?", p.WismaID).Error
	return &room, err
}

func (r RoomRepository) GetByUser(p *dto.UserRoomPaginateRequest) (*[]entity.Room, error) {
	var room []entity.Room
	offset := (p.Page - 1) * p.Limit
	err := r.db.Joins("Wisma").Joins("Wisma.User").Model(&entity.Room{}).Limit(p.Limit).Offset(offset).Find(&room, "\"Wisma\".\"user_id\" = ?", p.UserID.String()).Error
	return &room, err
}

func (r RoomRepository) Counter() (int64, error) {
	var count int64
	err := r.db.Model(&entity.Room{}).Count(&count).Error
	return count, err
}

func (r RoomRepository) Update(room *entity.Room) (*entity.Room, error) {
	err := r.db.Updates(room).Error
	return room, err
}

func (r RoomRepository) Delete(id string) error {
	var room entity.Room
	err := r.db.Where("id = ?", id).Delete(&room).Error
	return err
}
