package event

import (
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// TriggerStore contains basic operations for event trigger object.
type TriggerStore interface {
	InsertTrigger(Trigger) error
	FetchTrigger(gulid.ID, gulid.ID) (gulid.ID, error)
	ListTrigger(gulid.ID) ([]Trigger, error)
	RemoveTrigger(gulid.ID, gulid.ID) error
}
