package models

import "github.com/google/uuid"

//User model
type User struct {
	//'binding' is tag from GIN
	Id       uuid.UUID `json:"-" db:"id"`
	Name     string    `json:"name" binding:"required"`
	Username string    `json:"username" binding:"required"`
	Password string    `json:"password" binding:"required"`
}
