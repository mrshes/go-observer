package app

import (
	"context"
	"first-project/internal/config"
	"first-project/internal/database/migration"
	"first-project/internal/database/repository"
	"first-project/internal/router"
	"first-project/internal/server"
	"first-project/internal/services"
	"first-project/pkg/e"
	"first-project/pkg/gorm"
	"first-project/pkg/hash"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	gorm2 "gorm.io/gorm"
	"log"
)

type App struct {
	Log *logrus.Logger
}

func New(log *logrus.Logger) *App {
	return &App{
		Log: log,
	}
}

func Run(envFile string, configDir string) {
	// Load Env file configuration
	ctx := context.Background()
	err := godotenv.Load("./" + envFile)
	_ = e.New()
	if err != nil {
		log.Fatalln("Can't load env file", err)
	}
	// Init and fill configuration fields
	cfg, err := config.Init(configDir)
	if err != nil {
		logrus.Error(err)
	}

	// Init all services
	db, err := gorm.NewPostgres(postgres.Config{DSN: cfg.Postgres.ToString()}, &gorm2.Config{})
	migration.Run(db)
	if err != nil {
		log.Fatalln(err)

	}
	repositories := repository.NewRepositories(db)

	hasher := hash.NewSHA1Hasher(cfg.AuthSalt)

	services := services.NewServices(services.Deps{
		Ctx:    ctx,
		Repos:  repositories,
		Hasher: hasher,
		Cfg:    cfg,
	})

	// Init server
	router, err := router.NewRouter(services).Init()
	server := server.NewServer(cfg, router)
	err = server.Run()
	if err != nil {
		e.Log().Fatalln(err)
		panic(err)
	}
}
