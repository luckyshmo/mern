package pg

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/luckyshmo/api-example/models"
)

type UserPG struct {
	db *sqlx.DB
}

func NewUserPG(db *sqlx.DB) *UserPG {
	return &UserPG{db: db}
}

func (r *UserPG) GetAll() ([]models.User, error) {
	var userList []models.User

	query := fmt.Sprintf("SELECT name, username, id FROM %s", usersTable)
	err := r.db.Select(&userList, query)

	return userList, err
}

func (r *UserPG) GetById(userId uuid.UUID) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT name, username FROM %s WHERE id = $1", usersTable)
	err := r.db.Get(&user, query, userId)

	return user, err
}
