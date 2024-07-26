package logging_library

import (
	"github.com/neiasit/logging-library/core"
	"go.uber.org/fx"
)

var module = fx.Module(
	"logger",
	fx.Provide(
		core.LoadConfig,
		core.Logger,
	),
)
