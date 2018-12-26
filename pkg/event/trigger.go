package event

import gulid "github.com/elojah/game_01/pkg/ulid"

// TriggerStore set and get triggers.
type TriggerStore interface {
	SetTrigger(Trigger) error
	GetTrigger(gulid.ID, gulid.ID) (gulid.ID, error)
	ListTrigger(gulid.ID) ([]gulid.ID, error)
}
