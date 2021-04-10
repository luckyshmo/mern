package service

import (
	"github.com/google/uuid"
	"github.com/luckyshmo/api-example/models"
	"github.com/luckyshmo/api-example/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetById(userId uuid.UUID) (models.User, error) {
	return s.repo.GetById(userId)
}
