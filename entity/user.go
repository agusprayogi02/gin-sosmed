package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GenderType string

const (
	PRIA   GenderType = "pria"
	WANITA GenderType = "wanita"
)

func (me GenderType) String() string {
	switch me {
	case PRIA:
		return "pria"
	case WANITA:
		return "wanita"
	default:
		return "pria"
	}
}

type User struct {
	ID       uuid.UUID  `gorm:"type:varchar(60);primaryKey"`
	Name     string     `gorm:"type:varchar(150)"`
	Email    string     `gorm:"type:varchar(100);unique_index"`
	Password string     `gorm:"type:varchar(150)"`
	Gender   GenderType `sql:"type:ENUM('pria', 'wanita')" gorm:"type:varchar(8)"`
	gorm.Model
}
