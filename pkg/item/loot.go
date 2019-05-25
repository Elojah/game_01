package item

import gulid "github.com/elojah/game_01/pkg/ulid"

// LootStore interfaces operation to know if an entity is lootable.
type LootStore interface {
	FetchLoot(gulid.ID) (bool, error)
	UpsertLoot(gulid.ID) error
	RemoveLoot(gulid.ID) error
}
