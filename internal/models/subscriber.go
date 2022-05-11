package models

// Subscription ...
type Subscription struct {
	// Unique label of subscription
	Label string `json:"label"`
	// Blockchain net, `mainnet`, `rinkeby`, etc.
	Network string `json:"network"`
	// Address of Gnossis Safe (without prefix, just address)
	SafeAddress string `json:"safe_address"`
	// Unix of last event of Safe
	LastUpdated int64 `json:"last_updated"`
}
