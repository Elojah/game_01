package dto

import (
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
)

// NewACK convert a infra.ACK into a storage ACK.
func NewACK(ack infra.ACK) *ACK {
	return &ACK{
		ID: [16]byte(ack.ID),
	}
}

// Domain converts a storage ACK into a game ACK.
func (ack ACK) Domain() infra.ACK {
	return infra.ACK{
		ID: ulid.ID(ack.ID),
	}
}
