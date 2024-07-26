package core

import (
	"log/slog"
	"os"
)

func Logger(cfg *Config) *slog.Logger {
	var handler slog.Handler
	switch cfg.MODE {
	case devMode:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})
		break
	default:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
		break
	}
	logger := slog.New(handler)
	slog.SetDefault(logger)
	return logger
}
