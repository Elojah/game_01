package dto

import "github.com/elojah/game_01/pkg/ulid"

// SetPC represents the payload to send to create a new PC.
type SetPC struct {
	Token ulid.ID
	Type  ulid.ID
}

// ListPC represents the payload to list token PCs.
type ListPC struct {
	Token ulid.ID
}

// ConnectPC represents the payload to connect to an existing PC.
type ConnectPC struct {
	Token  ulid.ID
	Target ulid.ID
}

// DisconnectPC represents the payload to disconnect a token.
type DisconnectPC struct {
	Token ulid.ID
}

// EntityPC represents the response when connecting to an existing PC.
type EntityPC struct {
	ID ulid.ID
}