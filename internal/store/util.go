package store

import (
	"bytes"
	"encoding/gob"
)

// FromGob converts gob bytes to structure
func FromGob[T any](data []byte) (T, error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	value := new(T)
	if err := dec.Decode(value); err != nil {
		return *value, err
	}

	return *value, nil
}

// ToGob converts structure to gob bytes
func ToGob[T any](value T) ([]byte, error) {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(value); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
