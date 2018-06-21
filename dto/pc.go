package dto

import (
	game "github.com/elojah/game_01"
)

// SetPC represents the payload to send to create a new PC.
type SetPC struct {
	Token game.ID
	Type  game.ID
}

// ListPC represents the payload to list token PCs.
type ListPC struct {
	Token game.ID
}

// ConnectPC represents the payload to connect to an existing PC.
type ConnectPC struct {
	Token  game.ID
	Target game.ID
}

// Entity represents the response when connecting to an existing PC.
type Entity struct {
	ID game.ID
}
