package bot

import (
	"fmt"
	"strings"

	"github.com/MoonSHRD/dao-tg/pkg/gnosis"
)

// markdown escaper
var mdEscaper = strings.NewReplacer(
	`!`, `\!`,
	`*`, `\*`,
	`.`, `\.`,
	`-`, `\-`,
	`+`, `\+`,
	`=`, `\=`,
	`_`, `\_`,
	`~`, `\~`,
	"`", "\\`",
	`|`, `\|`,
	`[`, `\[`,
	`]`, `\]`,
	`(`, `\(`,
	`)`, `\)`,
	`<`, `\<`,
	`>`, `\>`,
	`{`, `\{`,
	`}`, `\}`,
	`|`, `\|`,
)

// dead code, but I'll leave it for now
func get2link(network gnosis.Network, address string) string {
	switch network {
	case gnosis.EthereumMainnet:
		return fmt.Sprintf("https://etherscan.io/address/%s", address)
	case gnosis.GnosisChain:
		return fmt.Sprintf("https://blockscout.com/xdai/mainnet/address/%s", address)
	case gnosis.Arbitrum:
		return fmt.Sprintf("https://arbiscan.io/address/%s", address)
	case gnosis.Avalanche:
		return fmt.Sprintf("https://snowtrace.io/address/%s", address)
	case gnosis.Aurora:
		return fmt.Sprintf("https://aurorascan.dev/address/%s", address)
	case gnosis.BinanceSmartChain:
		return fmt.Sprintf("https://bscscan.com/address/%s", address)
	case gnosis.EnergyWebChain:
		return fmt.Sprintf("https://explorer.energyweb.org/address/%s", address)
	case gnosis.Optimism:
		return fmt.Sprintf("https://optimistic.etherscan.io/address/%s", address)
	case gnosis.Polygon:
		return fmt.Sprintf("https://polygonscan.com/address/%s", address)
	case gnosis.Rinkeby:
		return fmt.Sprintf("https://rinkeby.etherscan.io/address/%s", address)
	case gnosis.Goerli:
		return fmt.Sprintf("https://goerli.etherscan.io/address/%s", address)
	case gnosis.Volta:
		return fmt.Sprintf("https://volta-explorer.energyweb.org/address/%s", address)
	default:
		return ""
	}
}
