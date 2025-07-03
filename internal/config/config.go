package config

import (
	"context"
	"time"

	"github.com/wallanaq/oauth2-token-introspection/internal/env"
)

type Config struct {
	Http   HttpConfig
	Logger LoggerConfig
	Server ServerConfig
}

type HttpConfig struct {
	Port int
}

type LoggerConfig struct {
	Type  string
	Level string
}

type ServerConfig struct {
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

func Load(ctx context.Context) (*Config, error) {

	var cfg Config

	cfg.Http.Port = env.GetInt("HTTP_PORT", 8080)
	cfg.Logger.Type = env.GetString("LOG_TYPE", "text")
	cfg.Logger.Level = env.GetString("LOG_LEVEL", "info")
	cfg.Server.ReadTimeout = env.GetDuration("SERVER_READ_TIMEOUT", 10*time.Second)
	cfg.Server.WriteTimeout = env.GetDuration("SERVER_WRITE_TIMEOUT", 10*time.Second)
	cfg.Server.IdleTimeout = env.GetDuration("SERVER_IDLE_TIMEOUT", 120*time.Second)
	cfg.Server.ShutdownTimeout = env.GetDuration("SERVER_SHUTDOWN_TIMEOUT", 5*time.Second)

	return &cfg, nil

}
