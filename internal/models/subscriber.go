package models

// Subscription ...
type Subscription struct {
	Network     string `json:"network"`
	SafeAddress string `json:"safe_address"`
	LastUpdated int64  `json:"last_updated"`
}
