package models

// EntityType ...
type EntityType = uint8

// EntityTypes ...
const (
	EntityUser EntityType = iota + 1
	EntityChat
	EntityChannel
)

// SubscriberFeed ...
type SubscriberFeed struct {
	EntityType  EntityType
	ID          int64
	AccessHash  int64
	LastUpdated int64
}
