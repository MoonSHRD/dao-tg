package gnosis

import "fmt"

// ExplorerLink ...
func ExplorerLink(network Network, address string) string {
	switch network {
	case EthereumMainnet:
		return fmt.Sprintf("https://etherscan.io/address/%s", address)
	case GnosisChain:
		return fmt.Sprintf("https://blockscout.com/xdai/mainnet/address/%s", address)
	case Arbitrum:
		return fmt.Sprintf("https://arbiscan.io/address/%s", address)
	case Avalanche:
		return fmt.Sprintf("https://snowtrace.io/address/%s", address)
	case Aurora:
		return fmt.Sprintf("https://aurorascan.dev/address/%s", address)
	case BinanceSmartChain:
		return fmt.Sprintf("https://bscscan.com/address/%s", address)
	case EnergyWebChain:
		return fmt.Sprintf("https://explorer.energyweb.org/address/%s", address)
	case Optimism:
		return fmt.Sprintf("https://optimistic.etherscan.io/address/%s", address)
	case Polygon:
		return fmt.Sprintf("https://polygonscan.com/address/%s", address)
	case Rinkeby:
		return fmt.Sprintf("https://rinkeby.etherscan.io/address/%s", address)
	case Goerli:
		return fmt.Sprintf("https://goerli.etherscan.io/address/%s", address)
	case Volta:
		return fmt.Sprintf("https://volta-explorer.energyweb.org/address/%s", address)
	default:
		return ""
	}
}
