package storage

import "errors"

var ErrNotFound = errors.New("not found")

type DB interface {
	// Get retrun key
	// Set set key value
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}
