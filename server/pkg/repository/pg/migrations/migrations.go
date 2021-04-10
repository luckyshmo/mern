package migrations

import (
	"errors"

	"github.com/luckyshmo/api-example/config"
	"github.com/sirupsen/logrus"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// runPgMigrations runs Postgres migrations
func RunPgMigrations() error { //? can be run from Makefile
	cfg := config.Get()
	if cfg.PgMigrationsPath == "" {
		logrus.Warn("No migration path provided")
		return nil
	}
	if cfg.PgHOST == "" || cfg.PgPORT == "" {
		return errors.New("No cfg.PgURL provided")
	}

	connectionString := "postgres://" + cfg.PgUserName + ":" + cfg.PgPAS + "@" + cfg.PgHOST + "/" + cfg.PgDBName + "?sslmode=" + cfg.PgSSLMode
	logrus.Info(connectionString)

	m, err := migrate.New(
		cfg.PgMigrationsPath,
		connectionString,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
