package dto

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type PostResponse struct {
	ID        uuid.UUID `json:"id"`
	Tweet     string    `json:"tweet"`
	Photo     *string   `json:"photo"`
	AuthorId  uuid.UUID `json:"author_id"`
	Author    User      `gorm:"foreignKey:AuthorId" json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostRequest struct {
	Tweet    string                `form:"tweet" binding:"required"`
	Photo    *multipart.FileHeader `form:"photo"`
	AuthorId uuid.UUID             `form:"author_id"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
