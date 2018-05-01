package game

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid"
)

// ID is an alias of ulid.ULID.
type ID = ulid.ULID

// NewULID returns a new random ID.
func NewULID() ID {
	return ID(ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader))
}
