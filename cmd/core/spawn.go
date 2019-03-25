package main

import (
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// Spawn moves the entity to spawn position and restore HP and MP.
func (a *app) Spawn(id gulid.ID, e event.E) error {

	sp := e.Action.Spawn
	ts := e.ID.Time()

	// #Retrieve previous target state.
	target, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrap(err, "spawn")
	}

	// #Check entity is dead
	if !target.Dead {
		return errors.Wrap(gerrors.ErrIsNotDead{EntityID: id.String()}, "spawn")
	}

	s, err := a.EntitySpawnStore.GetSpawn(sp.ID)
	if err != nil {
		return errors.Wrap(err, "spawn")
	}

	// #Move target at spawn position
	target.Position = s.Position
	// #Restore HP & MP
	target.HP = target.MaxHP
	target.MP = target.MaxMP

	// #Set entity new state.
	return errors.Wrap(a.EntityStore.SetEntity(target, ts), "spawn")
}
