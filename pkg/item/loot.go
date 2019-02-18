package item

import gulid "github.com/elojah/game_01/pkg/ulid"

// LootStore interfaces operation to know if an entity is lootable.
type LootStore interface {
	GetLoot(gulid.ID) (bool, error)
	SetLoot(gulid.ID) error
	DelLoot(gulid.ID) error
}
