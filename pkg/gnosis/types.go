package gnosis

import (
	"time"
)

type PaginatedResult[T any] struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []T     `json:"results"`
}

type Confirmation struct {
	Owner           string    `json:"owner"`
	SubmissionDate  time.Time `json:"submissionDate"`
	TransactionHash *string   `json:"transactionHash"`
	Signature       string    `json:"signature"`
	SignatureType   string    `json:"signatureType"`
}

type MultisigTransaction struct {
	Safe                  string         `json:"safe"`
	To                    string         `json:"to"`
	Value                 string         `json:"value"`
	Data                  *string        `json:"data"`
	Operation             int64          `json:"operation"`
	GasToken              string         `json:"gasToken"`
	SafeTxGas             int64          `json:"safeTxGas"`
	BaseGas               int64          `json:"baseGas"`
	GasPrice              string         `json:"gasPrice"`
	RefundReceiver        string         `json:"refundReceiver"`
	Nonce                 int64          `json:"nonce"`
	ExecutionDate         *time.Time     `json:"executionDate"`
	SubmissionDate        time.Time      `json:"submissionDate"`
	Modified              time.Time      `json:"modified"`
	BlockNumber           *int64         `json:"blockNumber"`
	TransactionHash       *string        `json:"transactionHash,omitempty"`
	SafeTxHash            string         `json:"safeTxHash,omitempty"`
	Executor              *string        `json:"executor,omitempty"`
	IsExecuted            bool           `json:"isExecuted,omitempty"`
	IsSuccessful          *bool          `json:"isSuccessful,omitempty"`
	EthGasPrice           *string        `json:"ethGasPrice,omitempty"`
	MaxFeePerGas          *string        `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas  *string        `json:"maxPriorityFeePerGas,omitempty"`
	GasUsed               *int64         `json:"gasUsed,omitempty"`
	Fee                   *string        `json:"fee,omitempty"`
	Origin                *string        `json:"origin,omitempty"`
	DataDecoded           interface{}    `json:"dataDecoded,omitempty"`
	ConfirmationsRequired *int64         `json:"confirmationsRequired,omitempty"`
	Confirmations         []Confirmation `json:"confirmations,omitempty"`
	Trusted               bool           `json:"trusted,omitempty"`
	Signatures            *string        `json:"signatures,omitempty"`
}

type TransferType string

const (
	TransferTypeEther TransferType = "ETHER_TRANSFER"
)

type Transfer struct {
	Type            TransferType `json:"type"`
	ExecutionDate   time.Time    `json:"executionDate"`
	BlockNumber     int64        `json:"blockNumber"`
	TransactionHash string       `json:"transactionHash"`
	To              string       `json:"to"`
	Value           string       `json:"value"`
	TokenID         *string      `json:"tokenId"`
	TokenAddress    *string      `json:"tokenAddress"`
	TokenInfo       *string      `json:"tokenInfo"`
	From            string       `json:"from"`
}

type TxType string

const (
	TxTypeEtheriumTransaction TxType = "ETHEREUM_TRANSACTION"
	TxTypeMultisigTransaction TxType = "MULTISIG_TRANSACTION"
)

type Transaction struct {
	Safe                  string         `json:"safe"`
	To                    string         `json:"to"`
	Value                 string         `json:"value"`
	Data                  *string        `json:"data"`
	Operation             int64          `json:"operation"`
	GasToken              string         `json:"gasToken"`
	SafeTxGas             int64          `json:"safeTxGas"`
	BaseGas               int64          `json:"baseGas"`
	GasPrice              string         `json:"gasPrice"`
	RefundReceiver        string         `json:"refundReceiver"`
	Nonce                 int64          `json:"nonce"`
	ExecutionDate         *time.Time     `json:"executionDate"`
	SubmissionDate        time.Time      `json:"submissionDate"`
	Modified              time.Time      `json:"modified"`
	BlockNumber           *int64         `json:"blockNumber"`
	TransactionHash       *string        `json:"transactionHash,omitempty"`
	SafeTxHash            string         `json:"safeTxHash,omitempty"`
	Executor              *string        `json:"executor,omitempty"`
	IsExecuted            bool           `json:"isExecuted,omitempty"`
	IsSuccessful          *bool          `json:"isSuccessful,omitempty"`
	EthGasPrice           *string        `json:"ethGasPrice,omitempty"`
	MaxFeePerGas          *string        `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas  *string        `json:"maxPriorityFeePerGas,omitempty"`
	GasUsed               *int64         `json:"gasUsed,omitempty"`
	Fee                   *string        `json:"fee,omitempty"`
	Origin                *string        `json:"origin,omitempty"`
	DataDecoded           interface{}    `json:"dataDecoded,omitempty"`
	ConfirmationsRequired *int64         `json:"confirmationsRequired,omitempty"`
	Confirmations         []Confirmation `json:"confirmations,omitempty"`
	Trusted               bool           `json:"trusted,omitempty"`
	Signatures            *string        `json:"signatures,omitempty"`
	Transfers             []Transfer     `json:"transfers"`
	TxType                TxType         `json:"txType"`
	TxHash                string         `json:"txHash,omitempty"`
	From                  string         `json:"from,omitempty"`
}

type SafeInfo struct {
	Address         string   `json:"address"`
	Nonce           int64    `json:"nonce"`
	Threshold       int64    `json:"threshold"`
	Owners          []string `json:"owners"`
	MasterCopy      string   `json:"masterCopy"`
	Modules         []string `json:"modules"`
	FallbackHandler string   `json:"fallbackHandler"`
	Guard           string   `json:"guard"`
	Version         string   `json:"version"`
}
