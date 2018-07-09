package dto

import (
	"errors"
	"strings"

	"github.com/elojah/game_01/pkg/ulid"
)

// SetPC represents the payload to send to create a new PC.
type SetPC struct {
	Token ulid.ID
	Name  string
	Type  ulid.ID
}

// Check checks setpc validity.
func (spc SetPC) Check() error {
	l := len(spc.Name)
	if l < 4 || l > 15 || strings.IndexFunc(spc.Name, func(r rune) bool {
		return r < 'A' || r > 'z'
	}) != -1 {
		return errors.New("invalid name")
	}
	return nil
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
