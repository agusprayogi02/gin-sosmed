package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID     uuid.UUID `gorm:"type:varchar(45);primaryKey"`
	Tweet  string    `gorm:"type:varchar(500)"`
	Photo  *string   `gorm:"type:varchar(150)"`
	UserID uuid.UUID
	User   User `gorm:"foreignKey:UserID"`
	gorm.Model
}
