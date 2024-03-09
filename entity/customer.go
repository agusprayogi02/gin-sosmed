package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID       uuid.UUID `gorm:"type:varchar(45);primaryKey"`
	Nik      string    `gorm:"type:varchar(16);unique"`
	Name     string    `gorm:"type:varchar(500)"`
	Address  *string   `gorm:"type:varchar(150)"`
	Phone    string    `gorm:"type:varchar(15)"`
	RoomID   uuid.UUID
	Room     Room `gorm:"foreignKey:RoomID"`
	UserID   uuid.UUID
	User     User `gorm:"foreignKey:UserID"`
	CheckIn  *time.Time
	CheckOut *time.Time
	gorm.Model
}
