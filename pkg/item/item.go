package item

import "github.com/elojah/game_01/pkg/ulid"

// Store is an interface for I object.
type Store interface {
	SetItem(I) error
	GetItem(ulid.ID) (I, error)
	DelItem(ulid.ID) error
}

// IsConsumable returns if an item is consumable for an entitty.
func (it I) IsConsumable() bool {
	switch it.Type.GetValue().(type) {
	case *Orb:
		return true
	}
	return false
}
