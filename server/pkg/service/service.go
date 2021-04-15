package service

import (
	"github.com/google/uuid"
	"github.com/luckyshmo/api-example/models"
	"github.com/luckyshmo/api-example/models/keep"
	"github.com/luckyshmo/api-example/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (uuid.UUID, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (uuid.UUID, error)
}

type User interface {
	GetById(userId uuid.UUID) (models.User, error)
	GetAll() ([]models.User, error)
}

type GoogleKeep interface {
	GetAll() (keep.Note, error)
}

type Service struct {
	Authorization
	User
	GoogleKeep
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
		GoogleKeep:    NewKeepService(),
	}
}
