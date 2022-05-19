package gnosis

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	"go.uber.org/fx"
)

var (
	ErrResponseNotOK = errors.New("status code is not 200")
)

var Module = fx.Module("gnosis", fx.Provide(
	New,
))

type Network string

const (
	EthereumMainnet   Network = "mainnet"
	GnosisChain       Network = "xdai"
	Arbitrum          Network = "arbitrum"
	Avalanche         Network = "avalanche"
	Aurora            Network = "aurora"
	BinanceSmartChain Network = "bsc"
	EnergyWebChain    Network = "ewc"
	Optimism          Network = "optiprism"
	Polygon           Network = "polygon"
	Goerli            Network = "goerli"
	Rinkeby           Network = "rinkeby"
	Volta             Network = "volta"
)

// NetworkFromPrefix creates Network from gnosis chain prefix
func NetworkFromPrefix(prefix string) (Network, error) {
	switch prefix {
	case "eth":
		return EthereumMainnet, nil
	case "gno":
		return GnosisChain, nil
	case "arbi":
		return Arbitrum, nil
	case "avax":
		return Avalanche, nil
	case "aurora":
		return Aurora, nil
	case "bnb":
		return BinanceSmartChain, nil
	case "ewt":
		return EnergyWebChain, nil
	case "oeth":
		return Optimism, nil
	case "matic":
		return Polygon, nil
	case "gor":
		return Goerli, nil
	case "rin":
		return Rinkeby, nil
	case "vt":
		return Volta, nil
	default:
		return EthereumMainnet, errors.New("no match with prefix")
	}
}

// AsPrefix ...
func (n Network) AsPrefix() string {
	switch n {
	case EthereumMainnet:
		return "eth"
	case GnosisChain:
		return "gno"
	case Arbitrum:
		return "arbi"
	case Avalanche:
		return "avax"
	case Aurora:
		return "aurora"
	case BinanceSmartChain:
		return "bnb"
	case EnergyWebChain:
		return "ewt"
	case Optimism:
		return "oeth"
	case Polygon:
		return "matic"
	case Goerli:
		return "gor"
	case Rinkeby:
		return "rin"
	case Volta:
		return "vt"
	default:
		return string(n)
	}
}

func (n Network) Label() string {
	switch n {
	case EthereumMainnet:
		return "Ethereum Mainnet"
	case GnosisChain:
		return "Gnosis Chain"
	case Arbitrum:
		return "Arbitrum"
	case Avalanche:
		return "Avalanche"
	case Aurora:
		return "Aurora"
	case BinanceSmartChain:
		return "Binance Smart Chain"
	case EnergyWebChain:
		return "Energy Web Chain"
	case Optimism:
		return "Optiprism"
	case Polygon:
		return "Polygon"
	case Goerli:
		return "Goerli"
	case Rinkeby:
		return "Rinkeby"
	case Volta:
		return "Volta"
	default:
		return string(n)
	}
}

func (n Network) String() string {
	return string(n)
}

// DefaultBase ...
const DefaultBase = "https://safe-transaction.gnosis.io/api/v1"

type Gnosis struct {
	Base   string
	Client *http.Client
}

func New() *Gnosis {
	return &Gnosis{
		Base:   DefaultBase,
		Client: http.DefaultClient,
	}
}

// WithNetwork creates new gnosis client with `network` and inherited http client.
// If `network` is incorrect it returns same client
func (g *Gnosis) WithNetwork(network Network) *Gnosis {
	if network != "" {
		return &Gnosis{
			Base:   fmt.Sprintf("https://safe-transaction.%s.gnosis.io/api/v1", network),
			Client: g.Client,
		}
	}

	return g
}

type PaginationOptions struct {
	// Which field to use ordering the results
	Ordering string `url:"ordering,omitempty"`
	Limit    *int64 `url:"limit,omitempty"`
	Offset   *int64 `url:"offset,omitempty"`
}

type AllTransactionsOptions struct {
	PaginationOptions
	// If `true`` only executed transactions are returned
	Executed *bool `url:"executed,omitempty"`
	// If `true` transactions with `nonce >= Safe current nonce` are also returned
	// Default: true
	Queued *bool `url:"queued,omitempty"`
	// If `true` just trusted transactions are shown (indexed, added by a delegate or with at least one confirmation)
	// Default: true
	Trusted *bool `url:"trusted,omitempty"`
}

type TransfersOptions struct {
	Limit  *int64 `url:"limit,omitempty"`
	Offset *int64 `url:"offset,omitempty"`

	BlockNumber   *int64 `url:"block_number,omitempty"`
	BlockNumberGt *int64 `url:"block_number__gt,omitempty"`
	BlockNumberLt *int64 `url:"block_number__lt,omitempty"`

	ExecutionDateGt  *string `url:"execution_date__gt,omitempty"`
	ExecutionDateLt  *string `url:"execution_date__lt,omitempty"`
	ExecutionDateGte *string `url:"execution_date__gte,omitempty"`
	ExecutionDateLte *string `url:"execution_date__lte,omitempty"`

	To *string `url:"to,omitempty"`

	TokenAddress    *string `url:"transaction_hash,omitempty"`
	TransactionHash *string `url:"transaction_hash,omitempty"`

	Value   *float64 `url:"value,omitempty"`
	ValueGt *float64 `url:"value__gt,omitempty"`
	ValueLt *float64 `url:"value__lt,omitempty"`

	ERC20  *string `url:"erc20,omitempty"`
	ERC721 *string `url:"erc721,omitempty"`
	Ether  *string `url:"ether,omitempty"`
}

type MultisigOptions struct {
	PaginationOptions
	Executed          *bool    `url:"executed,omitempty"`
	ExecutionDateGte  *string  `url:"execution_date__gte,omitempty"`
	ExecutionDateLte  *string  `url:"execution_date__lte,omitempty"`
	Failed            *string  `url:"failed,omitempty"`
	HasConfirmations  *string  `url:"has_confirmations,omitempty"`
	ModifiedGt        *string  `url:"modified__gt,omitempty"`
	ModifiedGte       *string  `url:"modified__gte,omitempty"`
	ModifiedLt        *string  `url:"modified__lt,omitempty"`
	ModifiedLte       *string  `url:"modified__lte,omitempty"`
	Nonce             *float64 `url:"nonce,omitempty"`
	NonceGt           *float64 `url:"nonce__gt,omitempty"`
	NonceGte          *float64 `url:"nonce__gte,omitempty"`
	NonceLt           *float64 `url:"nonce__lt,omitempty"`
	NonceLte          *float64 `url:"nonce__lte,omitempty"`
	Ordering          *string  `url:"ordering,omitempty"`
	SafeTxHash        *string  `url:"safe_tx_hash,omitempty"`
	SubmissionDateGte *string  `url:"submission_date__gte,omitempty"`
	SubmissionDateLte *string  `url:"submission_date__lte,omitempty"`
	To                *string  `url:"to,omitempty"`
	TransactionHash   *string  `url:"transaction_hash,omitempty"`
	Trusted           *string  `url:"trusted,omitempty"`
	Value             *float64 `url:"value,omitempty"`
	ValueGt           *float64 `url:"value__gt,omitempty"`
	ValueLt           *float64 `url:"value__lt,omitempty"`
}

func (g *Gnosis) Safe(safeAddress string) (*SafeInfo, error) {
	result := new(SafeInfo)
	url := fmt.Sprintf("%s/safes/%s", g.Base, safeAddress)
	if err := get(g.Client, url, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (g *Gnosis) IncomingTransfers(safeAddress string, options ...TransfersOptions) (*PaginatedResult[Transfer], error) {
	params, err := values(options...)
	if err != nil {
		return nil, err
	}

	result := new(PaginatedResult[Transfer])
	url := fmt.Sprintf("%s/safes/%s/incoming-transfers?%s", g.Base, safeAddress, params.Encode())
	if err = get(g.Client, url, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (g *Gnosis) Transactions(safeAddress string, options ...AllTransactionsOptions) (*PaginatedResult[Transaction], error) {
	params, err := values(options...)
	if err != nil {
		return nil, err
	}

	result := new(PaginatedResult[Transaction])
	url := fmt.Sprintf("%s/safes/%s/all-transactions?%s", g.Base, safeAddress, params.Encode())
	if err = get(g.Client, url, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (g *Gnosis) MultisigTransaction(safeTxHash string) (*MultisigTransaction, error) {
	result := new(MultisigTransaction)
	url := fmt.Sprintf("%s/multisig-transactions/%s", g.Base, safeTxHash)
	if err := get(g.Client, url, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (g *Gnosis) MultisigTransactionConfirmations(safeTxHash string, options ...PaginationOptions) (*PaginatedResult[Confirmation], error) {
	params, err := values(options...)
	if err != nil {
		return nil, err
	}

	result := new(PaginatedResult[Confirmation])
	url := fmt.Sprintf("%s/multisig-transactions/%s/confirmations?%s", g.Base, safeTxHash, params.Encode())
	if err = get(g.Client, url, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (g *Gnosis) MultisigTransactions(safeAddress string, options ...MultisigOptions) (*PaginatedResult[MultisigTransaction], error) {
	params, err := values(options...)
	if err != nil {
		return nil, err
	}

	result := new(PaginatedResult[MultisigTransaction])
	url := fmt.Sprintf("%s/safes/%s/multisig-transactions?%s", g.Base, safeAddress, params.Encode())
	if err = get(g.Client, url, result); err != nil {
		return nil, err
	}

	return result, nil
}

func get[T any](client *http.Client, url string, result *T) error {
	res, err := client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return ErrResponseNotOK
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, result)
}

func values[T any](options ...T) (url.Values, error) {
	if len(options) == 0 {
		return url.Values{}, nil
	}

	return query.Values(options[0])
}
