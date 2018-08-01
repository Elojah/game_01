package ulid

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid"
)

// ID is an alias of ulid.ULID.
type ID = ulid.ULID

// NewID returns a new random ID.
func NewID() ID {
	return ID(ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader))
}

// Parse follows internal oklog/ulid Parse.
func Parse(s string) (ID, error) {
	id, err := ulid.Parse(s)
	if err != nil {
		return ID{}, err
	}
	return ID(id), nil
}

// MustParse follows internal oklog/ulid MustParse.
func MustParse(s string) ID {
	return ID(ulid.MustParse(s))
}

// IsZero returns if id is zero.
// TODO move as method.
func IsZero(id ID) bool {
	return ulid.ULID(id).Time() == 0
}
