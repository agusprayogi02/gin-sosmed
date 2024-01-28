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
	Get(id string) (*dto.PostResponse, error)
	GetAll(p *dto.PaginateRequest) (*int64, *[]dto.PostResponse, error)
	Update(req *dto.PostEditRequest) (*dto.PostResponse, error)
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

func (s *postService) GetAll(p *dto.PaginateRequest) (*int64, *[]dto.PostResponse, error) {
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
		data = append(data, dto.PostResponse{
			ID:       post.ID,
			AuthorId: post.UserID,
			Tweet:    post.Tweet,
			Photo:    post.Photo,
			Author: dto.User{
				ID:    post.User.ID.String(),
				Name:  post.User.Name,
				Email: post.User.Email,
			},
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}
	return &count, &data, nil
}

func (s *postService) Update(req *dto.PostEditRequest) (*dto.PostResponse, error) {
	var data *dto.PostResponse

	err := s.repo.Update
}
