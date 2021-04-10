package config

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/kelseyhightower/envconfig"
)

// Config. Should be filled from Env. Use launch.json(vscode) on local machine
type Config struct {
	PgHOST           string `envconfig:"PG_HOST"`
	PgPORT           string `envconfig:"PG_PORT"`
	PgPAS            string `envconfig:"PG_PAS"`
	PgSSLMode        string `envconfig:"PG_SSLMODE"`
	PgMigrationsPath string `envconfig:"PG_MIGRATIONS_PATH"`
	PgUserName       string `envconfig:"PG_USERNAME"`
	PgDBName         string `envconfig:"PG_DBNAME"`

	Environment string `envconfig:"ENV"`

	AppPort  string `envconfig:"APP_PORT"`
	LogLevel string `envconfig:"LOG_LEVEL"`
}

var (
	config Config
	once   sync.Once
)

// Get reads config from environment. Once.
func Get() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			logrus.Fatal(err)
		}
		validate(config)
	})
	return &config
}

func validate(cfg Config) { //TODO? logging isn't configure at this moment... probably return message or error?
	refConf := reflect.ValueOf(cfg)
	typeOfRefConf := refConf.Type()

	for i := 0; i < refConf.NumField(); i++ {
		if fmt.Sprint(refConf.Field(i).Interface()) == "" {
			logrus.Warn(fmt.Sprintf("Config: %s value is empty!", typeOfRefConf.Field(i).Name))
		}
	}
}
