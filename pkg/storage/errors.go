package storage

import (
	"errors"
)

var (
	// ErrNotFound is raised when a mandatory resource is not found in storage.
	ErrNotFound = errors.New("no results found")
)
