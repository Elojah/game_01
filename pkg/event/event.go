package event

import (
	"github.com/elojah/game_01/pkg/infra"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// QStore contains basic queue operations for event E object.
type QStore interface {
	PublishEvent(E, gulid.ID) error
	SubscribeEvent(gulid.ID) *infra.Subscription
}

// Store contains basic operations for event E object.
type Store interface {
	Insert(E, gulid.ID) error
	Fetch(gulid.ID, gulid.ID) (E, error)
	List(gulid.ID, gulid.ID) ([]E, error)
	Remove(gulid.ID, gulid.ID) error
}

// App contains event stores and applications.
type App interface {
	QStore
	Store
	TriggerStore

	Create(E, gulid.ID) error
	Cancel(E) error
}
