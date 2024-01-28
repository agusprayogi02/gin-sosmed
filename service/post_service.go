package service

import (
	"fmt"
	"os"
	"path/filepath"

	"gin-sosmed/config"
	"gin-sosmed/dto"
	"gin-sosmed/entity"
	"gin-sosmed/errorhandler"
	"gin-sosmed/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PostService interface {
	Create(req *dto.PostRequest) error
	Get(id string) (*dto.PostResponse, error)
	GetAll(p *dto.PaginateRequest, host string) (*int64, *[]dto.PostResponse, error)
	Update(c *gin.Context) (*dto.PostResponse, error)
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
		UserID: req.AuthorId,
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

func (p *postService) Get(id string) (*dto.PostResponse, error) {
	var post *dto.PostResponse

	data, err := p.repo.Get(id)
	if err != nil {
		return nil, &errorhandler.NotFoundError{
			Message: err.Error(),
		}
	}
	post = &dto.PostResponse{
		ID:       data.ID,
		Tweet:    data.Tweet,
		Photo:    data.Photo,
		AuthorId: data.UserID,
		Author: dto.User{
			ID:    data.User.ID.String(),
			Name:  data.User.Name,
			Email: data.User.Email,
		},
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return post, nil
}

func (s *postService) GetAll(p *dto.PaginateRequest, host string) (*int64, *[]dto.PostResponse, error) {
	var data []dto.PostResponse

	posts, err := s.repo.GetAll(p)
	if err != nil {
		return nil, nil, &errorhandler.NotFoundError{
			Message: err.Error(),
		}
	}
	count, err := s.repo.Counter()
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}
	for _, post := range *posts {
		pst := dto.PostResponse{
			ID:       post.ID,
			AuthorId: post.UserID,
			Tweet:    post.Tweet,
			Author: dto.User{
				ID:    post.User.ID.String(),
				Name:  post.User.Name,
				Email: post.User.Email,
			},
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}
		if post.Photo != nil {
			tmp := fmt.Sprintf("http://%v/%v", host, *post.Photo)
			pst.Photo = &tmp
		}
		data = append(data, pst)
	}
	return &count, &data, nil
}

func (s *postService) Update(c *gin.Context) (*dto.PostResponse, error) {
	var req dto.PostEditRequest

	if err := c.ShouldBind(&req); err != nil {
		return nil, &errorhandler.UnprocessableEntityError{
			Message: err.Error(),
		}
	}

	id := c.Param("id")
	oldPost, err := s.repo.Get(id)
	if err != nil {
		return nil, &errorhandler.NotFoundError{
			Message: err.Error(),
		}
	}
	oldPost.Tweet = req.Tweet

	if req.Photo != nil {
		if err := os.MkdirAll(config.TweetsFolder, 0o755); err != nil {
			return nil, &errorhandler.InternalServerError{
				Message: err.Error(),
			}
		}
		ext := filepath.Ext(req.Photo.Filename)
		newFileName := filepath.Join(config.TweetsFolder, uuid.New().String()+ext)
		if err := c.SaveUploadedFile(req.Photo, newFileName); err != nil {
			return nil, &errorhandler.InternalServerError{
				Message: err.Error(),
			}
		}
		req.Photo.Filename = newFileName
		oldPost.Photo = &newFileName
	}

	updatedPost, err := s.repo.Update(oldPost)
	if err != nil {
		return nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}
	data := dto.PostResponse{
		ID:       updatedPost.ID,
		Tweet:    updatedPost.Tweet,
		Photo:    updatedPost.Photo,
		AuthorId: updatedPost.UserID,
		Author: dto.User{
			ID:    updatedPost.User.ID.String(),
			Name:  updatedPost.User.Name,
			Email: updatedPost.User.Email,
		},
		CreatedAt: updatedPost.CreatedAt,
		UpdatedAt: updatedPost.UpdatedAt,
	}
	return &data, nil
}
