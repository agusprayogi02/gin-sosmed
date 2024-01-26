package service

import (
	"gin-sosmed/dto"
	"gin-sosmed/entity"
	"gin-sosmed/errorhandler"
	"gin-sosmed/repository"

	"github.com/google/uuid"
)

type PostService interface {
	Create(req *dto.PostRequest) error
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(p repository.PostRepository) *postService {
	return &postService{
		repo: p,
	}
}

func (p *postService) Create(req *dto.PostRequest) error {
	post := entity.Post{
		ID:     uuid.New(),
		UserID: *req.AuthorId,
		Tweet:  req.Tweet,
	}

	if req.Photo != nil {
		post.Photo = &req.Photo.Filename
	}

	if err := p.repo.Create(&post); err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return nil
}
