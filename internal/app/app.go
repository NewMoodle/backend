package app

import (
	"context"
	"errors"
	"github.com/ZhansultanS/myLMS/backend/internal/api"
	"github.com/ZhansultanS/myLMS/backend/internal/config"
	"github.com/ZhansultanS/myLMS/backend/internal/repository"
	"github.com/ZhansultanS/myLMS/backend/internal/server"
	"github.com/ZhansultanS/myLMS/backend/internal/service"
	"github.com/ZhansultanS/myLMS/backend/pkg/database"
	"github.com/ZhansultanS/myLMS/backend/pkg/hasher"
	"github.com/ZhansultanS/myLMS/backend/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start(env string) {
	var err error

	log := logger.NewLogrus(env)
	cfg, err := config.Load(env)
	if err != nil {
		log.Error(err.Error())
		return
	}

	pool, err := database.NewPostgreConnectionPool(
		cfg.Postgre.Host, cfg.Postgre.Port, cfg.Postgre.Username, cfg.Postgre.Password, cfg.Postgre.Database,
		cfg.Postgre.SSLMode, cfg.Postgre.Pools,
	)
	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Info("Database pool connection established")
	defer pool.Close()

	repositories := repository.NewRepository(pool)
	passwordhasher := hasher.BcryptHasher{Cost: cfg.Auth.PasswordHashCost}

	deps := service.Deps{
		Repositories:   repositories,
		PasswordHasher: &passwordhasher,
	}
	services := service.NewService(deps)

	handlers := api.New(log, services)
	srv := server.New(cfg.Http, handlers.Init(cfg.GinMode))

	go func() {
		if err = srv.Start(); !errors.Is(err, http.ErrServerClosed) {
			log.Error(err.Error())
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	if err = srv.Stop(ctx); err != nil {
		log.Error(err.Error())
	}
}
