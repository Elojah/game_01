package entity

import "github.com/elojah/game_01/pkg/ulid"

const (
	DefaultInventorySize = 42
)

// InventoryStore contains basic operations for entity inventory object.
type InventoryStore interface {
	FetchInventory(ulid.ID) (Inventory, error)
	UpsertInventory(Inventory) error
	RemoveInventory(ulid.ID) error
}

// MRInventoryStore contains basic operations for entity most recent inventory object.
type MRInventoryStore interface {
	FetchMRInventory(ulid.ID) (Inventory, error)
	UpsertMRInventory(ulid.ID, Inventory) error
	RemoveMRInventory(ulid.ID) error
}
