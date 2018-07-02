package event

import "github.com/elojah/game_01/pkg/ulid"

// Recurrer requires the receiver to create a new recurrer with subject ID.
type Recurrer struct {
	ID       ulid.ID
	EntityID ulid.ID
	TokenID  ulid.ID
	Action   QAction
}

// QRecurrerMapper handles send/receive methods for recurrers.
type QRecurrerMapper interface {
	PublishRecurrer(Recurrer, ulid.ID) error
	SubscribeRecurrer(ulid.ID) *Subscription
}
