package infra

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// QRecurrerStore contains basic queue operations for infra recurrer object.
type QRecurrerStore interface {
	PublishRecurrer(Recurrer, ulid.ID) error
	SubscribeRecurrer(ulid.ID) *Subscription
}

// CoreStore contains basic operations for infra recurrer object.
type RecurrerStore interface {
	UpsertRecurrer(Recurrer) error
	FetchRecurrer(ulid.ID) (Recurrer, error)
	RemoveRecurrer(ulid.ID) error
}

// RecurrerApp contains recurrer stores and applications.
type RecurrerApp interface {
	QRecurrerStore
	RecurrerStore
	SyncStore

	Create(ulid.ID, ulid.ID) (Recurrer, error)
	Erase(ulid.ID) error
}
