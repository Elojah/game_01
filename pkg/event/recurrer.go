package event

import "github.com/elojah/game_01/pkg/ulid"

// Recurrer requires the receiver to create a new recurrer with subject ID.
type Recurrer struct {
	TokenID  ulid.ID
	EntityID ulid.ID
	Action   QAction
	Pool     ulid.ID
}

// QRecurrerMapper handles send/receive methods for recurrers.
type QRecurrerMapper interface {
	PublishRecurrer(Recurrer, ulid.ID) error
	SubscribeRecurrer(ulid.ID) *Subscription
}

// RecurrerMapper handles set/get methods for recurrers.
type RecurrerMapper interface {
	SetRecurrer(Recurrer) error
	GetRecurrer(RecurrerSubset) (Recurrer, error)
	DelRecurrer(RecurrerSubset) error
}

// RecurrerSubset retrieves recurrer by Token ID.
type RecurrerSubset struct {
	TokenID ulid.ID
}
