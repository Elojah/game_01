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

// PCStore contains basic operations fo entity PC object.
type PCStore interface {
	InsertPC(PC, ulid.ID) error
	FetchPC(ulid.ID, ulid.ID) (PC, error)
	ListPC(ulid.ID) ([]PC, error)
	RemovePC(ulid.ID, ulid.ID) error
}

// PCLeft represents the number of character an account can still create.
type PCLeft int

// PCLeftStore contains basic operations fo entity PCLeft object.
type PCLeftStore interface {
	InsertPCLeft(PCLeft, ulid.ID) error
	FetchPCLeft(ulid.ID) (PCLeft, error)
	RemovePCLeft(ulid.ID) error
}
