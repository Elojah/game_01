package item

import "github.com/elojah/game_01/pkg/ulid"

// Store is an interface for I object.
type Store interface {
	SetItem(I) error
	GetItem(ulid.ID) (I, error)
	DelItem(ulid.ID) error
}
