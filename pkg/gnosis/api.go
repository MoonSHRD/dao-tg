package gnosis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Base API Domain
var Base = "https://safe-client.gnosis.io/v1"

// GetTransaction ...
func GetTransaction(client *http.Client, transactionType string, safeAddress string, safeTxHash string) (*DetailedTransaction, *http.Response, error) {
	res, err := client.Get(fmt.Sprintf("%s/chains/4/transactions/%s_%s_%s", Base, transactionType, safeAddress, safeTxHash))
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res, err
	}

	multisig := new(DetailedTransaction)
	if err := json.Unmarshal(data, &multisig); err != nil {
		return nil, res, err
	}

	return multisig, res, nil
}

// GetHistory ...
func GetHistory(client *http.Client, safeAddress string) ([]Transaction, *http.Response, error) {
	res, err := client.Get(fmt.Sprintf("%s/chains/4/safes/%s/transactions/history", Base, safeAddress))
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res, err
	}

	apiResult := new(ApiResult)
	if err := json.Unmarshal(data, apiResult); err != nil {
		return nil, res, err
	}

	transactions := make([]Transaction, 0)
	for _, result := range apiResult.Results {
		if result.Type != EntryTypeTransaction {
			continue
		}

		transactions = append(transactions, result.Value.(Transaction))
	}

	return transactions, res, nil
}

// GetQueue ...
func GetQueue(client *http.Client, safeAddress string) ([]Transaction, *http.Response, error) {
	res, err := client.Get(fmt.Sprintf("%s/chains/4/safes/%s/transactions/queued", Base, safeAddress))
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res, err
	}

	apiResult := new(ApiResult)
	if err := json.Unmarshal(data, apiResult); err != nil {
		return nil, res, err
	}

	transactions := make([]Transaction, 0)
	for _, result := range apiResult.Results {
		if result.Type != EntryTypeTransaction {
			continue
		}

		transactions = append(transactions, result.Value.(Transaction))
	}

	return nil, res, nil
}
