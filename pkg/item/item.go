package item

import "github.com/elojah/game_01/pkg/ulid"

// Store contains basic operations for item I object.
type Store interface {
	Upsert(I) error
	Fetch(ulid.ID) (I, error)
	Remove(ulid.ID) error
}

// App contains items stores and applications.
type App interface {
	Store
	LootStore
}

// IsConsumable returns if an item is consumable for an entitty.
func (it I) IsConsumable() bool {
	switch it.Type.GetValue().(type) {
	case *Orb:
		return true
	}
	return false
}
