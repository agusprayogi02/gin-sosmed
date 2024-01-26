package repository

import (
	"gin-sosmed/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(p *entity.Post) error
	Get(id string) (*entity.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}

func (p *postRepository) Create(req *entity.Post) error {
	return p.db.Create(req).Error
}

func (p *postRepository) Get(id string) (*entity.Post, error) {
	var post *entity.Post
	err := p.db.Preload("User").First(&post, "id = ?", id).Error
	return post, err
}
