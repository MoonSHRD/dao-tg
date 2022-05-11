package models

import "time"

// Subscription ...
type Subscription struct {
	// Unique label of subscription
	Label string `json:"label"`
	// Blockchain net, `mainnet`, `rinkeby`, etc.
	Network string `json:"network"`
	// Address of Gnossis Safe (without prefix, just address)
	SafeAddress string `json:"safe_address"`
	// Time when last event happened of Safe
	LastUpdated time.Time `json:"last_updated"`
}
