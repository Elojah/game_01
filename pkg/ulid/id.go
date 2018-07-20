package ulid

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid"
)

// ID is an alias of ulid.ULID.
type ID = [16]byte

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

// String convert a ulid into a string.
func String(id ID) string {
	return ulid.ULID(id).String()
}

// Compare compare two ids.
func Compare(lhs ID, rhs ID) int {
	return ulid.ULID(lhs).Compare(ulid.ULID(rhs))
}

// IsZero returns if id is zero.
func IsZero(id ID) bool {
	return ulid.ULID(id).Time() == 0
}
