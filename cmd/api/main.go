package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/wallanaq/oauth2-token-introspection/internal/config"
	"github.com/wallanaq/oauth2-token-introspection/internal/handler"
	"github.com/wallanaq/oauth2-token-introspection/internal/logger"
	"github.com/wallanaq/oauth2-token-introspection/internal/server"
)

func main() {

	ctx := context.Background()

	cfg, err := config.Load(ctx)
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	logger := logger.New(cfg)

	if err := run(ctx, cfg, logger); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

}

func run(ctx context.Context, cfg *config.Config, logger *slog.Logger) error {

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()
	mux.Handle("/health/", http.StripPrefix("/health", handler.NewHealthHandler()))
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", handler.NewIntrospectionHandler()))

	srv := server.NewHTTPServer(cfg, mux)

	errCh := make(chan error, 1)

	logger.Info("starting server...", slog.String("addr", srv.Addr))

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- fmt.Errorf("listen and serve: %w", err)
		}
	}()

	select {
	case <-ctx.Done():

		logger.Info("shutdown signal received, stopping server...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("shutdown server: %w", err)
		}

	case err := <-errCh:
		return err
	}

	return nil

}
