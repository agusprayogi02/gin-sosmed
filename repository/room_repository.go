package repository

import (
	"time"

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
	err := r.db.Joins("Wisma").Joins("Wisma.User").Model(&entity.Room{}).First(&room, "rooms.id = ?", id).Error
	if err != nil {
		return room, err
	}
	return room, nil
}

func (r RoomRepository) GetAll(p *dto.PaginateRequest) (*[]entity.Room, error) {
	var room []entity.Room
	offset := (p.Page - 1) * p.Limit
	err := r.db.Joins("Wisma").Joins("Wisma.User").Model(&entity.Room{}).Limit(p.Limit).Offset(offset).Find(&room).Error
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
	err := r.db.Joins("Wisma").Joins("Wisma.User").Model(&entity.Room{}).Limit(p.Limit).Offset(offset).Find(&room, "Wisma.user_id = ?", p.UserID.String()).Error
	return &room, err
}

func (r RoomRepository) GetByUserRaw(p *dto.UserRoomPaginateRequest) (*[]dto.RoomResponse, error) {
	var rooms *[]dto.RoomResponse
	offset := (p.Page - 1) * p.Limit
	err := r.db.Table("rooms").Joins("LEFT JOIN wismas Wisma ON rooms.wisma_id = Wisma.id AND Wisma.deleted_at IS NULL").
		Joins("LEFT JOIN users Wisma__User ON Wisma.user_id = Wisma__User.id AND Wisma__User.deleted_at IS NULL").
		Select("rooms.*, Wisma.id AS Wisma__id,Wisma.user_id AS Wisma__user_id,Wisma.name AS Wisma__name,Wisma.address AS Wisma__address,Wisma.code AS Wisma__code,Wisma.note AS Wisma__note,Wisma.created_at AS Wisma__created_at,Wisma.updated_at AS Wisma__updated_at,Wisma.deleted_at AS Wisma__deleted_at,Wisma__User.id AS Wisma__User__id,Wisma__User.name AS Wisma__User__name,Wisma__User.email AS Wisma__User__email,Wisma__User.password AS Wisma__User__password,Wisma__User.gender AS Wisma__User__gender,Wisma__User.role AS Wisma__User__role,Wisma__User.created_at AS Wisma__User__created_at,Wisma__User.updated_at AS Wisma__User__updated_at,Wisma__User.deleted_at AS Wisma__User__deleted_at, IFNULL((SELECT 'terisi' FROM customers WHERE customers.room_id = rooms.id AND customers.check_in <= ? AND customers.check_out >= ? LIMIT 1), 'kosong') AS status", time.Now(), time.Now()).
		Limit(p.Limit).Offset(offset).Where("Wisma.user_id = ? AND rooms.deleted_at IS NULL", p.UserID.String()).Scan(&rooms).Error

	return rooms, err
}

func (r RoomRepository) Counter() (int64, error) {
	var count int64
	err := r.db.Model(&entity.Room{}).Count(&count).Error
	return count, err
}

func (r RoomRepository) Update(room *entity.Room) error {
	return r.db.Updates(room).Error
}

func (r RoomRepository) Delete(id string) error {
	var room entity.Room
	err := r.db.Where("id = ?", id).Delete(&room).Error
	return err
}
