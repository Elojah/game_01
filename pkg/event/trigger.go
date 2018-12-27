package event

import (
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// TriggerStore set and get triggers.
type TriggerStore interface {
	SetTrigger(Trigger) error
	GetTrigger(gulid.ID, gulid.ID) (gulid.ID, error)
	ListTrigger(gulid.ID) ([]Trigger, error)
	DelTrigger(gulid.ID, gulid.ID) error
}

// TriggerService handles trigger event interactions.
type TriggerService interface {
	Set(E, gulid.ID) error
}
