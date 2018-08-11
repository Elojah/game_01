package ulid

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid"
)

// ID is an alias of ulid.ULID.
type ID ulid.ULID

// NewID returns a new random ID.
func NewID() ID {
	return ID(ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader))
}

// IsZero returns if id is zero.
func (id ID) IsZero() bool {
	return ulid.ULID(id).Time() == 0
}

// Bytes returns id as byte slice for protobuf marshaling.
func (id ID) Bytes() []byte {
	return id[:]
}

// Marshal never returns any error..
func (id ID) Marshal() ([]byte, error) {
	return id.Bytes(), nil
}

// MarshalTo never returns any error.
func (id ID) MarshalTo(data []byte) (n int, err error) {
	copy(data[0:16], id[:])
	return 16, nil
}

// Unmarshal never returns any error.
func (id *ID) Unmarshal(data []byte) error {
	var tmp [16]byte
	copy(data[0:16], tmp[:])
	*id = tmp
	return nil
}

// Size always returns 16.
func (id *ID) Size() int {
	return 16
}

// MarshalJSON returns id in human readable string format.
func (id ID) MarshalJSON() ([]byte, error) {
	return ulid.ULID(id).MarshalText()
}

// UnmarshalJSON unmarshals and valid data.
func (id *ID) UnmarshalJSON(data []byte) error {
	var u ulid.ULID
	if err := u.UnmarshalText(data); err != nil {
		return err
	}
	*id = ID(u)
	return nil
}

// only required if the compare option is set
func (id ID) Compare(other ID) int {
	return id.Compare(other)
}

// only required if the equal option is set
func (id ID) Equal(other ID) bool {
	return id.Equal(other)
}

// only required if populate option is set
func NewPopulatedID(r randyID) *ID {
	id := NewID()
	return &id
}

type randyID interface{}
