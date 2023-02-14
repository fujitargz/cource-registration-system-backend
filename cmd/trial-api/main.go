package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"

	"github.com/fujitargz/cource-registration-system-backend/internal/infra"
	"github.com/fujitargz/cource-registration-system-backend/internal/interface/handler"
	"github.com/fujitargz/cource-registration-system-backend/internal/interface/router"
	"github.com/fujitargz/cource-registration-system-backend/internal/interface/server"
	"github.com/fujitargz/cource-registration-system-backend/internal/usecase"
)

func main() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	db, err := infra.Open()
	if err != nil {
		return 1
	}
	defer db.Close()

	userRepository := infra.NewUserRepositoryInfra(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)
	appHandler := handler.NewAppHandler(userHandler)
	router := router.NewRouter(appHandler)
	s := server.NewServer(8000, router)

	errCh := make(chan error, 1)
	termCh := make(chan os.Signal, 1)
	signal.Notify(termCh, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		errCh <- s.Start()
	}()

	select {
	case <-termCh:
		s.Stop(ctx)
		return 0
	case <-errCh:
		return 1
	}
}
