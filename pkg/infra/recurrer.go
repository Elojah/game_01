package infra

import "github.com/elojah/game_01/pkg/ulid"

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
