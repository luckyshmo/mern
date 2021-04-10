package pg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/luckyshmo/api-example/pkg/repository/pg/migrations"
)

const (
	usersTable = "user_tb"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = migrations.RunPgMigrations() //TODO config? rename package to pgMigration?
	if err != nil {
		return nil, err
	}

	return db, nil
}
