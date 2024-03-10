package dto

import (
	"time"

	"github.com/google/uuid"
)

type CustomerResponse struct {
	ID       uuid.UUID    `json:"id"`
	Nik      string       `json:"nik"`
	Name     string       `json:"name"`
	Address  *string      `json:"address"`
	Phone    string       `json:"phone"`
	RoomID   uuid.UUID    `json:"room_id"`
	Room     RoomResponse `gorm:"foreignKey:RoomID" json:"room"`
	UserID   uuid.UUID    `json:"user_id"`
	User     User         `gorm:"foreignKey:UserID" json:"user"`
	CheckIn  *time.Time   `json:"check_in"`
	CheckOut *time.Time   `json:"check_out"`
}

type CustomerRequest struct {
	Nik      string    `json:"nik" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Address  *string   `json:"address" binding:"required"`
	Phone    string    `json:"phone" binding:"required"`
	RoomID   string    `json:"room_id" binding:"required"`
	UserID   uuid.UUID `json:"user_id" binding:"required"`
	CheckIn  time.Time `json:"check_in" binding:"required"`
	CheckOut time.Time `json:"check_out" binding:"required"`
}
