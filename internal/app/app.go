package app

import (
	"avitoTestTask/internal/config"
	"avitoTestTask/internal/core/models"
	"avitoTestTask/internal/infrastructure/repository/postgres"
	"avitoTestTask/internal/infrastructure/transport/http/handlers"
	"avitoTestTask/internal/infrastructure/transport/http/server"
	"avitoTestTask/internal/usecases"
	"fmt"

	"github.com/sirupsen/logrus"
	p "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func RunServer() {
	cfg, err := config.InitConfig()
	if err != nil {
		logrus.Fatalf("failed to init config: %v", err)
	}

	db, err := initPostgresDB(&cfg.DB)
	if err != nil {
		logrus.Fatalf("failed to init DB: %v", err)
	}

	if err := db.AutoMigrate(&models.Team{}, &models.User{}, &models.PullRequest{}); err != nil {
		logrus.Fatalf("failed to run DB migrate: %v", err)
	}

	repo := postgres.NewPostgresRepo(db)
	usecases := usecases.NewUseCases(repo.UserRepo, repo.TeamRepo, repo.PullRequestRepo, repo.StatsRepo)
	handlers := handlers.NewHandlers(usecases)

	httpServer := server.NewServer(handlers, &cfg.Server)
	if err := httpServer.Run(); err != nil {
		logrus.Fatalf("server error: %v", err)
	}
}

func initPostgresDB(cfg *config.DB) (*gorm.DB, error) {
	dsn := makeDSN(cfg)
	db, err := gorm.Open(p.Open(dsn))
	return db, err
}

func makeDSN(cfg *config.DB) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC", cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)
}
