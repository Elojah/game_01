package main

import (
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

func (a *app) LootSource(id ulid.ID, e event.E) error {

	loot := e.Action.GetValue().(*event.LootSource)
	ts := e.ID.Time()

	// #Retrieve entity
	source, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrapf(err, "get entity %s", id.String())
	}

	_, _ = loot, source
	return nil
}
