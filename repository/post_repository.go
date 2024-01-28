package repository

import (
	"gin-sosmed/dto"
	"gin-sosmed/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(p *entity.Post) error
	Get(id string) (*entity.Post, error)
	GetAll(p *dto.PaginateRequest) (*[]entity.Post, error)
	Counter() (int64, error)
	Update(p *entity.Post) (*entity.Post, error)
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

func (r *postRepository) Get(id string) (*entity.Post, error) {
	var post *entity.Post
	err := r.db.Preload("User").First(&post, "id = ?", id).Error
	return post, err
}

func (r *postRepository) GetAll(p *dto.PaginateRequest) (*[]entity.Post, error) {
	var posts *[]entity.Post
	offset := (p.Page - 1) * p.Limit
	err := r.db.Preload("User").Limit(p.Limit).Offset(offset).Find(posts).Error
	return posts, err
}

func (r *postRepository) Counter() (int64, error) {
	var count int64
	err := r.db.Count(&count).Error
	return count, err
}

func (r *postRepository) Update(post *entity.Post) (*entity.Post, error) {
	err := r.db.Preload("User").Updates(post).Error
	return post, err
}
