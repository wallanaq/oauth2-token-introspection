package logger

import (
	"log/slog"
	"os"

	"github.com/wallanaq/oauth2-token-introspection/internal/config"
)

var levels = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

func New(config *config.Config) *slog.Logger {

	logLevel, ok := levels[config.Logger.Level]
	if !ok {
		logLevel = slog.LevelInfo
	}

	var handler slog.Handler

	switch config.Logger.Type {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})
	case "text":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})
	default:
		handler = slog.Default().Handler()
	}

	logger := slog.New(handler)

	slog.SetDefault(logger)

	return logger

}
