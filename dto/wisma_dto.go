package dto

import (
	"time"

	"github.com/google/uuid"
)

type WismaResponse struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	Nama      string    `json:"nama"`
	Address   string    `json:"address"`
	Code      string    `json:"code"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WismaEditRequest struct {
	UserID  *uuid.UUID `json:"user_id"`
	Nama    string     `json:"nama"`
	Address string     `json:"address"`
	Code    string     `json:"code"`
	Note    string     `json:"note"`
}

type WismaRequest struct {
	UserID  uuid.UUID `json:"user_id"`
	Nama    string    `json:"nama" binding:"required"`
	Address string    `json:"address" binding:"required"`
	Code    string    `json:"code" binding:"required"`
	Note    string    `json:"note"`
}
