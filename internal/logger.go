package internal

import (
	"log/slog"
	"os"
)

func Logger(cfg *Config) *slog.Logger {
	switch cfg.MODE {
	case prodMode:
		return slog.New(slog.NewJSONHandler(os.Stdout, nil))
	case devMode:
		return slog.New(slog.NewTextHandler(os.Stdout, nil))
	default:
		return slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}
}
