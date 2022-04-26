package contract

import (
	"github.com/MoonSHRD/dao-tg/config"
	"github.com/MoonSHRD/dao-tg/contract/gnosis"
	"github.com/ethereum/go-ethereum/ethclient"
)

// NewEthClient ...
func NewEthClient(cfg config.Ethereum) (*ethclient.Client, error) {
	return ethclient.Dial(cfg.Gateway)
}

// NewGnosis ...
func NewGnosis(client *ethclient.Client, cfg config.Gnosis) (*gnosis.Gnosis, error) {
	return gnosis.NewGnosis(cfg.ContractAddress, client)
}
