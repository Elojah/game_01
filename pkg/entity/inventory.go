package entity

import "github.com/elojah/game_01/pkg/ulid"

const (
	DefaultInventorySize = 42
)

// InventoryStore wraps inventory interactions.
type InventoryStore interface {
	GetInventory(ulid.ID) (Inventory, error)
	SetInventory(Inventory) error
	DelInventory(ulid.ID) error
}

// MRInventoryStore stores Most Recent inventory for each entity.
type MRInventoryStore interface {
	GetMRInventory(ulid.ID) (Inventory, error)
	SetMRInventory(ulid.ID, Inventory) error
	DelMRInventory(ulid.ID) error
}

// InventoryService wraps operations for MR logic.
type InventoryService interface {
	Get(ulid.ID, ulid.ID) (Inventory, error)
	SetMR(ulid.ID, Inventory) error
}
