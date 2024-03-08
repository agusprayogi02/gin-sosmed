package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Room struct {
	ID      uuid.UUID `gorm:"type:varchar(45);primaryKey"`
	WismaID *uuid.UUID
	Wisma   *Wisma `gorm:"foreignKey:WismaID"`
	Name    string `gorm:"type:varchar(10)"`
	Note    string `gorm:"type:varchar(255)"`
	gorm.Model
}
