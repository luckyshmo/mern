package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/luckyshmo/api-example/models"
	"github.com/luckyshmo/api-example/pkg/repository/pg"
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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: pg.NewAuthPostgres(db),
		User:          pg.NewUserPG(db),
	}
}
