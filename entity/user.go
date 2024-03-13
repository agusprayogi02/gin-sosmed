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

type RoleType string

const (
	ADMIN RoleType = "admin"
	USER  RoleType = "user"
)

func (me RoleType) String() string {
	switch me {
	case ADMIN:
		return "admin"
	case USER:
		return "user"
	default:
		return "user"
	}
}

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
	ID       uuid.UUID  `gorm:"type:varchar(45);primaryKey"`
	Name     string     `gorm:"type:varchar(150)"`
	Email    string     `gorm:"type:varchar(100);unique_index"`
	Password string     `gorm:"type:varchar(150)"`
	Gender   GenderType `gorm:"type:ENUM('pria', 'wanita')"`
	Role     RoleType   `gorm:"type:ENUM('admin', 'user');default:'user'"`
	gorm.Model
}
