package entity

import (
	"github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/ulid"
)

const (
	// MaxPC is the maximum number of characters an account can have.
	MaxPC = 4
)

var (
	pcNames = map[string]struct{}{
		"trickster":  {},
		"mesmerist":  {},
		"inquisitor": {},
		"totemist":   {},
		"scavenger":  {},
	}
)

// PC alias an entity.
type PC = E

// Check checks if pc fields are valid.
func (pc PC) Check() error {
	if _, ok := pcNames[pc.Name]; !ok {
		return errors.ErrInvalidEntityType{EntityType: pc.Name}
	}
	return nil
}

// PCStore is an interface to create a new PC.
type PCStore interface {
	SetPC(PC, ulid.ID) error
	GetPC(ulid.ID, ulid.ID) (PC, error)
	ListPC(ulid.ID) ([]PC, error)
	DelPC(ulid.ID, ulid.ID) error
}

// PCLeft represents the number of character an account can still create.
type PCLeft int

// PCLeftStore interfaces creation/retrieval of PCLeft.
type PCLeftStore interface {
	SetPCLeft(PCLeft, ulid.ID) error
	GetPCLeft(ulid.ID) (PCLeft, error)
	DelPCLeft(ulid.ID) error
}
