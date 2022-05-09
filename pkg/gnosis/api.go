package gnosis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	"go.uber.org/fx"
)

var Module = fx.Module("gnosis", fx.Provide(
	NewWithConfig,
))

// DefaultBase ...
const DefaultBase = "https://safe-transaction.gnosis.io/api/v1"

type Config struct {
	Base   string
	Client *http.Client
}

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

func NewWithConfig(config Config) *Gnosis {
	return &Gnosis{
		Base:   config.Base,
		Client: config.Client,
	}
}

type PaginationOptions struct {
	// Which field to use ordering the results
	Ordering string `url:"ordering,omitempty"`
	Limit    *int64 `url:"limit,omitempty"`
	Offset   *int64 `url:"offset,omitempty"`
}

type TransactionOptions struct {
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

type MultisigOptions struct {
	PaginationOptions
	Executed          *bool    `url:"executed,omitempty"`
	ExecutionDateGte  *string  `url:"execution_date__gte,omitempty"`
	ExecutionDateLte  *string  `url:"execution_date__lte,omitempty"`
	Failed            *string  `url:"failed,omitempty"`
	HasConfirmations  *string  `url:"has_confirmations,omitempty"`
	ModifiedGt        *string  `url:"modified_gt,omitempty"`
	ModifiedGte       *string  `url:"modified_gte,omitempty"`
	ModifiedLt        *string  `url:"modified_lt,omitempty"`
	ModifiedLte       *string  `url:"modified_lte,omitempty"`
	Nonce             *float64 `url:"nonce,omitempty"`
	NonceGt           *float64 `url:"nonce_gt,omitempty"`
	NonceGte          *float64 `url:"nonce_gte,omitempty"`
	NonceLt           *float64 `url:"nonce_lt,omitempty"`
	NonceLte          *float64 `url:"nonce_lte,omitempty"`
	Ordering          *string  `url:"ordering,omitempty"`
	SafeTxHash        *string  `url:"safe_tx_hash,omitempty"`
	SubmissionDateGte *string  `url:"submission_date_gte,omitempty"`
	SubmissionDateLte *string  `url:"submission_date_lte,omitempty"`
	To                *string  `url:"to,omitempty"`
	TransactionHash   *string  `url:"transaction_hash,omitempty"`
	Trusted           *string  `url:"trusted,omitempty"`
	Value             *float64 `url:"value,omitempty"`
	ValueGt           *float64 `url:"value_gt,omitempty"`
	ValueLt           *float64 `url:"value_lt,omitempty"`
}

func (g *Gnosis) Safe(safeAddress string) (*SafeInfo, error) {
	result := new(SafeInfo)
	url := fmt.Sprintf("%s/safes/%s", g.Base, safeAddress)
	if err := get(g.Client, url, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (g *Gnosis) Transactions(safeAddress string, options ...TransactionOptions) (*PaginatedResult[Transaction], error) {
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

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, result); err != nil {
		return err
	}

	return err
}

func values[T any](options ...T) (url.Values, error) {
	if len(options) == 0 {
		return url.Values{}, nil
	}

	return query.Values(options[0])
}
