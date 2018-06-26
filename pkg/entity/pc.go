package entity

import (
	game "github.com/elojah/game_01"
)

const (
	// MaxPC is the maximum number of characters an account can have.
	MaxPC = 4
)

var (
	pcNames = map[string]struct{}{
		"trickster":  struct{}{},
		"mesmerist":  struct{}{},
		"inquisitor": struct{}{},
		"totemist":   struct{}{},
		"scavenger":  struct{}{},
	}
)

// PC alias an entity.
type PC E

// Check checks if pc fields are valid.
func (pc PC) Check() error {
	if _, ok := pcNames[pc.Name]; !ok {
		return game.ErrInvalidEntityType
	}
	return nil
}

// PCMapper is an interface to create a new PC.
type PCMapper interface {
	SetPC(PC, game.ID) error
	GetPC(PCSubset) (PC, error)
	ListPC(PCSubset) ([]PC, error)
}

// PCSubset represents a subset of PC by account ID.
type PCSubset struct {
	ID        game.ID
	AccountID game.ID
}

// PCLeft represents the number of character an account can still create.
type PCLeft int

// PCLeftMapper interfaces creation/retrieval of PCLeft.
type PCLeftMapper interface {
	SetPCLeft(PCLeft, game.ID) error
	GetPCLeft(PCLeftSubset) (PCLeft, error)
}

// PCLeftSubset represents a subset of PCLeft per account.
type PCLeftSubset struct {
	AccountID game.ID
}
