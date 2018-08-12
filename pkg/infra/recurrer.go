package infra

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// QRecurrerStore handles send/receive methods for recurrers.
type QRecurrerStore interface {
	PublishRecurrer(Recurrer, ulid.ID) error
	SubscribeRecurrer(ulid.ID) *Subscription
}

// RecurrerStore handles set/get methods for recurrers.
type RecurrerStore interface {
	SetRecurrer(Recurrer) error
	GetRecurrer(RecurrerSubset) (Recurrer, error)
	DelRecurrer(RecurrerSubset) error
}

// RecurrerSubset retrieves recurrer by Token ID.
type RecurrerSubset struct {
	TokenID ulid.ID
}

// RecurrerService represents recurrer usecases.
type RecurrerService interface {
	New(ulid.ID, ulid.ID) (Recurrer, error)
	Remove(ulid.ID) error
}
