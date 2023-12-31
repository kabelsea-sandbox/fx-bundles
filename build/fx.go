package build

import (
	"go.uber.org/fx"
)

var ModuleName = "build"

var Module = fx.Module(
	ModuleName,

	fx.Provide(
		NewContext,
	),
)
