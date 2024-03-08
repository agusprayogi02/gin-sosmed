package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Wisma struct {
	ID      uuid.UUID `gorm:"type:varchar(45);primaryKey"`
	UserID  *uuid.UUID
	User    *User  `gorm:"foreignKey:UserID"`
	Name    string `gorm:"type:varchar(100)"`
	Address string `gorm:"type:varchar(200)"`
	Code    string `gorm:"type:varchar(100)"`
	Note    string `gorm:"type:varchar(255)"`
	gorm.Model
}
