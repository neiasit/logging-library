package pkg

import (
	"github.com/neiasit/logging-library/internal"
	"go.uber.org/fx"
)

var module = fx.Module(
	"logger",
	fx.Provide(
		internal.LoadConfig,
		internal.Logger,
	),
)
