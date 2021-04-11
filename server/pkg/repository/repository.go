package repository

import (
	"github.com/google/uuid"
	"github.com/luckyshmo/api-example/models"
	"github.com/luckyshmo/api-example/pkg/repository/mongo"
)

type Authorization interface {
	CreateUser(user models.User) (uuid.UUID, error)
	GetUser(username, passwordHash string) (models.User, error)
}

type User interface {
	GetById(userId uuid.UUID) (models.User, error)
	GetAll() ([]models.User, error)
}

type Repository struct {
	Authorization
	User
}

func NewRepository(mg *mongo.MongoClient) *Repository {
	return &Repository{
		Authorization: mongo.NewAuthMongo(mg),
		User:          mongo.NewUserMongo(mg),
	}
}
