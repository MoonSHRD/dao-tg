package store

import (
	"bytes"

	"github.com/akrylysov/pogreb"
)

// FilterItemIterator is an iterator over DB key-value pairs. It iterates the items with prefix in an unspecified order.
type FilterItemIterator struct {
	prefix []byte
	iter   *pogreb.ItemIterator
}

// Next returns the next filtered key-value pair if available, otherwise it returns ErrIterationDone error.
func (i FilterItemIterator) Next() ([]byte, []byte, error) {
	for {
		key, val, err := i.iter.Next()
		if err != nil {
			return nil, nil, err
		}

		if bytes.HasPrefix(key, i.prefix) {
			return key, val, nil
		}
	}
}
