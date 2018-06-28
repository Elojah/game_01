package infra

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// ACK represents an validated packet ID
type ACK struct {
	ID ulid.ID
}
