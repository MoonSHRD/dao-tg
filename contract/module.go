package contract

import (
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Module("blockchain", fx.Provide(
	NewEthClient,
	NewGnosis,
))
