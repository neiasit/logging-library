package logging_lib

import (
	core2 "github.com/neiasit/logging-lib/core"
	"go.uber.org/fx"
)

var module = fx.Module(
	"logger",
	fx.Provide(
		core2.LoadConfig,
		core2.Logger,
	),
)
