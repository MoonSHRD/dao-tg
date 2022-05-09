package gnosis

import (
	"github.com/MoonSHRD/dao-tg/pkg/gnosis/client"
	"github.com/go-openapi/strfmt"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Module("gnosis",
	fx.Supply(
		client.DefaultTransportConfig().WithHost("safe-transaction.rinkeby.gnosis.io"),
	),
	fx.Provide(
		strfmt.NewFormats,
		client.NewHTTPClientWithConfig,
	),
)
