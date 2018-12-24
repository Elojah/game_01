package entity

import "github.com/elojah/game_01/pkg/ulid"

// InventoryStore wraps inventory interactions.
type InventoryStore interface {
	GetInventory(ulid.ID) (Inventory, error)
	SetInventory(Inventory) error
}
