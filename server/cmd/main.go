package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/luckyshmo/api-example/config"
	"github.com/luckyshmo/api-example/pkg/external"
	"github.com/luckyshmo/api-example/pkg/external/pythonNotes"
	"github.com/luckyshmo/api-example/pkg/handler"
	"github.com/luckyshmo/api-example/pkg/repository"
	"github.com/luckyshmo/api-example/pkg/repository/mongo"
	"github.com/luckyshmo/api-example/pkg/repository/pg"
	"github.com/luckyshmo/api-example/pkg/service"
	"github.com/luckyshmo/api-example/server"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	//Mat Ryer advice to handle all app errors
	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

//main func
func run() error {
	// config
	cfg := config.Get() //? errors before logrus init

	// logger configuration
	lvl, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel) //using debug lvl if we can't parse
		logrus.Warn("Using debug level logger")
	} else {
		logrus.SetLevel(lvl)
	}
	if cfg.Environment == "production" {
		var JSONF = new(logrus.JSONFormatter)
		JSONF.TimestampFormat = time.RFC3339
		logrus.SetFormatter(JSONF)
	}

	//Init GRPC Clinet
	client, err := pythonNotes.NewKeepClient("localhost:50051")
	if err != nil {
		return errors.Wrap(err, "failed init keep grpc client")
	}

	go func() {
		ext := external.NewExternalSource(client)
		notes, err := ext.NoteGetter.GetAll()
		if err != nil {
			logrus.Error(err)
		}
		logrus.Println(notes)
	}()

	//Init PG
	if false {
		// _ "github.com/lib/pq"
		_, err := pg.NewPostgresDB(pg.Config{ //? you can get db by config
			Host:     cfg.PgHOST,
			Port:     cfg.PgPORT,
			Username: cfg.PgUserName,
			DBName:   cfg.PgDBName,
			SSLMode:  cfg.PgSSLMode,
			Password: cfg.PgPAS,
		})
		if err != nil {
			return errors.Wrap(err, "failed to initialize db")
		}
	}
	//Init MONGO
	mc, err := mongo.NewMongoClient(*cfg) //TODO return close function for graceful shutdown
	if err != nil {
		return errors.Wrap(err, "failed to initialize db")
	}

	//Init main components
	//Good Clean arch and dependency injection example
	repos := repository.NewRepository(mc)
	services := service.NewService(repos)

	// note, err := services.GoogleKeep.GetAll()
	// if err != nil {
	// 	return err
	// }
	// fmt.Print(note)
	// return nil

	handlers := handler.NewHandler(services)

	//starting server
	srv := new(server.Server) //TODO? server.Server should be *serviceName*.server
	go func() {
		if err := srv.Run(cfg.AppPort, handlers.InitRoutes()); err != nil {
			logrus.Error(fmt.Sprintf("error occured while running http server: %s", err.Error()))
		}
	}()

	logrus.Print("App Started")

	quit := make(chan os.Signal, 1)
	//if app get SIGTERM it will exit
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App Shutting Down")

	// if err := db.Close(); err != nil {
	// 	return err
	// }

	return nil
}
