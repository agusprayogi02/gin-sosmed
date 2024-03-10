package dto

import (
	"time"

	"github.com/google/uuid"
)

type RoomResponse struct {
	ID        uuid.UUID  `json:"id"`
	WismaID   *uuid.UUID `json:"wisma_id"`
	Wisma     *Wisma     `gorm:"foreignKey:WismaID" json:"wisma"`
	Name      string     `json:"name"`
	Capacity  int        `json:"capacity"`
	Note      string     `json:"note"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type RoomEditRequest struct {
	WismaID uuid.UUID `json:"wisma_id"`
	Name    string    `json:"name"`
	Note    string    `json:"note"`
}

type RoomRequest struct {
	WismaID  string `json:"wisma_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Capacity int    `json:"capacity"`
	Note     string `json:"note"`
}

type Wisma struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Code      string    `json:"code"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RoomPaginateRequest struct {
	Page    int    `form:"page"`
	Limit   int    `form:"limit"`
	WismaID string `form:"wisma_id" binding:"required"`
}
