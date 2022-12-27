package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fujitargz/cource-registration-system-backend/internal/http"
)

func Run() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	s := http.NewServer(8000)

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
