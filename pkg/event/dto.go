package event

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// DTO is a event object sent from client to server for an input/action.
type DTO struct {
	ID     ulid.ID
	Token  ulid.ID
	TS     int64
	Action interface{}
}
