package dto

import (
	"time"

	"github.com/google/uuid"
)

type RoomResponse struct {
	ID        uuid.UUID      `json:"id"`
	WismaID   *uuid.UUID     `json:"wisma_id"`
	Wisma     *WismaResponse `gorm:"foreignKey:WismaID" json:"wisma"`
	Name      string         `json:"name"`
	Capacity  int            `json:"capacity"`
	Note      string         `json:"note"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type RoomEditRequest struct {
	WismaID  uuid.UUID `json:"wisma_id"`
	Name     string    `json:"name"`
	Note     string    `json:"note"`
	Capacity int       `json:"capacity"`
}

type RoomRequest struct {
	WismaID  string `json:"wisma_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Capacity int    `json:"capacity" binding:"required"`
	Note     string `json:"note"`
}

type RoomPaginateRequest struct {
	Page    int    `form:"page"`
	Limit   int    `form:"limit"`
	WismaID string `form:"wisma_id" binding:"required"`
}
type UserRoomPaginateRequest struct {
	Page   int       `form:"page"`
	Limit  int       `form:"limit"`
	UserID uuid.UUID `form:"user_id"`
}
