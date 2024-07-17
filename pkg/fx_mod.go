package pkg

import (
	"github.com/neiasit/logging-library/internal"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log/slog"
)

var Module = fx.Module(
	"logger",
	fx.Provide(
		internal.LoadConfig,
		internal.Logger,
	),
	fx.WithLogger(func(logger *slog.Logger) fxevent.Logger {
		return &fxevent.SlogLogger{
			Logger: logger,
		}
	}),
)
