package gnosis

import (
	"encoding/json"
	"fmt"
	"time"
)

// Timestamp ...
type Timestamp struct {
	time.Time
}

// UnmarshalJSON decodes an int64 timestamp into a time.Time object
func (p *Timestamp) UnmarshalJSON(bytes []byte) error {
	// 1. Decode the bytes into an int64
	var raw int64
	err := json.Unmarshal(bytes, &raw)

	if err != nil {
		return fmt.Errorf("error decoding timestamp: %s", err)
	}

	// 2 - Parse the unix timestamp
	*&p.Time = time.Unix(raw/1000, raw%1000*1_000_000)
	return nil
}

// MarshalJSON encodes an time.Time timestamp into a json unix time number
func (p *Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.UnixMilli())
}

// TxStatus ...
type TxStatus = string

// TxStatus consts
const (
	StatusAwaitingConfirmations TxStatus = "AWAITING_CONFIRMATIONS"
	StatusSuccess               TxStatus = "SUCCESS"
)

// Address on blockchain
type Address struct {
	Value string `json:"value"`
}

// DetailedTransaction ...
type DetailedTransaction struct {
	SafeAddress           string                 `json:"safeAddress"`
	TxID                  string                 `json:"txId"`
	ExecutedAt            *Timestamp             `json:"executedAt"`
	TxStatus              TxStatus               `json:"txStatus"`
	TxInfo                TxInfo                 `json:"txInfo"`
	TxData                *TxData                `json:"txData"`
	DetailedExecutionInfo *DetailedExecutionInfo `json:"detailedExecutionInfo"`
	TxHash                *string                `json:"txHash"`
}

// Transaction ...
type Transaction struct {
	ID            string         `json:"id"`
	Timestamp     Timestamp      `json:"timestamp"`
	TxStatus      TxStatus       `json:"txStatus"`
	TxInfo        TxInfo         `json:"txInfo"`
	ExecutionInfo *ExecutionInfo `json:"executionInfo"`
}

// TxInfoType ...
type TxInfoType = string

// TxInfoType consts
const (
	InfoTypeCreation       TxInfoType = "Creation"
	InfoTypeTransfer       TxInfoType = "Transfer"
	InfoTypeSettingsChange TxInfoType = "SettingsChange"
)

// TxInfo ...
type TxInfo struct {
	Type TxInfoType  `json:"type"`
	Info interface{} `json:"-"`
}

// UnmarshalJSON ...
func (i *TxInfo) UnmarshalJSON(data []byte) error {
	var txType struct {
		Type TxInfoType `json:"type"`
	}

	if err := json.Unmarshal(data, &txType); err != nil {
		return err
	}

	i.Type = txType.Type
	switch i.Type {
	case InfoTypeTransfer:
		var inter Transfer
		if err := json.Unmarshal(data, &inter); err != nil {
			return err
		}
		i.Info = inter
	default:
		inter := map[string]interface{}{}
		if err := json.Unmarshal(data, &inter); err != nil {
			return err
		}
		i.Info = inter
	}

	return nil
}

// MarshalJSON ...
func (i *TxInfo) MarshalJSON() ([]byte, error) {
	switch i.Type {
	case InfoTypeTransfer:
		inter := struct {
			Type string "json:\"type\""
			Transfer
		}{
			Type:     i.Type,
			Transfer: i.Info.(Transfer),
		}
		return json.Marshal(&inter)
	default:
		inter := map[string]interface{}{}
		for k, v := range i.Info.(map[string]interface{}) {
			inter[k] = v
		}
		inter["type"] = i.Type
		return json.Marshal(&inter)
	}
}

// Transfer ...
type Transfer struct {
	Sender       Address      `json:"sender"`
	Recipient    Address      `json:"recipient"`
	Direction    string       `json:"direction"`
	TransferInfo TransferInfo `json:"transferInfo"`
}

// TransferInfo ...
type TransferInfo struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// TxData ...
type TxData struct {
	HexData     *string `json:"hexData"`
	DataEncoded *string `json:"dataEncoded"`
	To          Address `json:"to"`
	Value       string  `json:"value"`
	Operation   uint64  `json:"operation"`
}

// ExecutionInfo ...
type ExecutionInfo struct {
	Type                   string    `json:"type"`
	Nonce                  uint64    `json:"nonce"`
	ConfirmationsRequired  uint64    `json:"confirmationsRequired"`
	ConfirmationsSubmitted uint64    `json:"confirmationsSubmitted"`
	MissingSigners         []Address `json:"missingSigners"`
}

// DetailedExecutionInfo ...
type DetailedExecutionInfo struct {
	Type                  string         `json:"type"`
	SubmittedAt           Timestamp      `json:"submittedAt"`
	Nonce                 uint64         `json:"nonce"`
	SafeTxGas             string         `json:"safeTxGas"`
	BaseGas               string         `json:"baseGas"`
	GasPrice              string         `json:"gasPrice"`
	GasToken              string         `json:"gasToken"`
	RefundReciever        Address        `json:"refundReciever"`
	SafeTxHash            string         `json:"safeTxHash"`
	Executor              *Address       `json:"executor,omitempty"`
	Signers               []Address      `json:"signers"`
	ConfirmationsRequired uint64         `json:"confirmationsRequired"`
	Confirmations         []Confirmation `json:"confirmations"`
}

// Confirmation ...
type Confirmation struct {
	Signer      Address   `json:"signer"`
	Signature   string    `json:"signature"`
	SubmittedAt Timestamp `json:"submittedAt"`
}

type ApiResult struct {
	Results []ApiResultEntry `json:"results"`
}

type ApiResultEntryType = string

const (
	EntryTypeTransaction ApiResultEntryType = "TRANSACTION"
	EntryTypeDateLabel   ApiResultEntryType = "DATE_LABEL"
	EntryTypeLabel       ApiResultEntryType = "LABEL"
)

type ApiResultEntry struct {
	Type  ApiResultEntryType `json:"type"`
	Value interface{}        `json:"-"`
}

// UnmarshalJSON ...
func (i *ApiResultEntry) UnmarshalJSON(data []byte) error {
	var txType struct {
		Type TxInfoType `json:"type"`
	}

	if err := json.Unmarshal(data, &txType); err != nil {
		return err
	}

	i.Type = txType.Type
	switch i.Type {
	case EntryTypeTransaction:
		var inter struct {
			Transaction Transaction `json:"transaction"`
		}
		if err := json.Unmarshal(data, &inter); err != nil {
			return err
		}
		i.Value = inter.Transaction
	default:
		inter := map[string]interface{}{}
		if err := json.Unmarshal(data, &inter); err != nil {
			return err
		}
		i.Value = inter
	}

	return nil
}

// MarshalJSON ...
func (i *ApiResultEntry) MarshalJSON() ([]byte, error) {
	switch i.Type {
	case EntryTypeTransaction:
		var inter = struct {
			Type        string      `json:"type"`
			Transaction Transaction `json:"transaction"`
		}{
			Type:        i.Type,
			Transaction: i.Value.(Transaction),
		}
		return json.Marshal(&inter)
	default:
		info := map[string]interface{}{}
		for k, v := range i.Value.(map[string]interface{}) {
			info[k] = v
		}
		info["type"] = i.Type
		return json.Marshal(&info)
	}
}
