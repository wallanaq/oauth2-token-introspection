package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/wallanaq/oauth2-token-introspection/internal/config"
	"github.com/wallanaq/oauth2-token-introspection/internal/http"
)

func main() {

	ctx := context.Background()

	if err := run(ctx); err != nil {
		slog.Error("fatal error", "msg", err.Error())
		os.Exit(1)
	}

}

func run(ctx context.Context) error {

	cfg, err := config.Load(".", "..")
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := http.NewAPIServer(cfg)

	if err := srv.Start(ctx); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil

}
