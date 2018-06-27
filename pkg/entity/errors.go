package entity

import (
	"errors"
)

var (
	// ErrInvalidEntityType is raised when an entity doesn't respect the correct type.
	ErrInvalidEntityType = errors.New("invalid entity type")
)
