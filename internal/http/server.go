package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/wallanaq/oauth2-token-introspection/internal/config"
)

type APIServer struct {
	srv             *http.Server
	shutdownTimeout time.Duration
}

func NewAPIServer(cfg *config.Config) *APIServer {
	return &APIServer{
		srv: &http.Server{
			Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
			ReadTimeout:  cfg.Server.ReadTimeout,
			WriteTimeout: cfg.Server.WriteTimeout,
			IdleTimeout:  cfg.Server.IdleTimeout,
		},
		shutdownTimeout: cfg.Server.ShutdownTimeout,
	}
}

func (s *APIServer) Start(ctx context.Context) error {

	slog.Info("starting server...", "addr", s.srv.Addr)

	errCh := make(chan error, 1)

	go func() {
		defer close(errCh)

		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		slog.Error("failed to start server")
		return fmt.Errorf("failed to start server: %w", err)
	case <-ctx.Done():
		slog.Warn("received shutdown signal")
		return s.Shutdown(ctx)
	}

}

func (s *APIServer) Shutdown(ctx context.Context) error {

	slog.Info("shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(ctx, s.shutdownTimeout)
	defer cancel()

	if err := s.srv.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("failed to shutdown server: %w", err)
	}

	slog.Info("server shutdown gracefully")

	return nil

}
