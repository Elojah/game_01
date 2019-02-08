package ulid

import (
	"crypto/rand"
	"encoding/json"
	"time"

	"github.com/oklog/ulid"
)

// ID is an alias of ulid.ULID.
type ID ulid.ULID

// NewID returns a new random ID.
func NewID() ID {
	return ID(ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader))
}

// NewIDs returns a new array of random IDs.
func NewIDs(n int) []ID {
	t := ulid.Timestamp(time.Now())
	ids := make([]ID, n)
	for i := range ids {
		ids[i] = ID(ulid.MustNew(t, rand.Reader))
	}
	return ids
}

// NewTimeID returns a new random ID based on time t.
func NewTimeID(t uint64) ID {
	return ID(ulid.MustNew(t, rand.Reader))
}

// MustParse alias ulid.MustParse. Panics if s is invalid.
func MustParse(s string) ID {
	return ID(ulid.MustParse(s))
}

// Parse alias ulid.Parse. Panics if s is invalid.
func Parse(s string) (ID, error) {
	id, err := ulid.Parse(s)
	return ID(id), err
}

// Time returns ms time of ID.
func (id ID) Time() uint64 {
	return ulid.ULID(id).Time()
}

// IsZero returns if id is zero.
func (id ID) IsZero() bool {
	return id.Time() == 0
}

// Zero returns a zero id.
func Zero() ID {
	return ID{}
}

// String returns a human readable string ID.
func (id ID) String() string {
	return ulid.ULID(id).String()
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
	if len(data) != 16 {
		return ulid.ErrBufferSize
	}
	for i := 0; i < 16; i++ {
		id[i] = data[i]
	}
	return nil
}

// Size always returns 16.
func (id *ID) Size() int {
	return 16
}

// MarshalJSON returns id in human readable string format.
func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

// UnmarshalJSON unmarshals and valid data.
func (id *ID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	u, err := ulid.Parse(s)
	if err != nil {
		return err
	}
	*id = ID(u)
	return nil
}

// Compare only required if the compare option is set
func (id ID) Compare(other ID) int {
	return ulid.ULID(id).Compare(ulid.ULID(other))
}

// Equal only required if the equal option is set
func (id ID) Equal(other ID) bool {
	return id.Compare(other) == 0
}

// NewPopulatedID only required if populate option is set
func NewPopulatedID(r randyID) *ID {
	id := ID(ulid.MustNew(uint64(r.Uint32()), rand.Reader))
	return &id
}

type randyID interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}
