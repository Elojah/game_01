package app

import "github.com/elojah/game_01/pkg/item"

var _ item.App = (*A)(nil)

// A implementation of item applications.
type A struct {
	item.Store
	item.LootStore
}
