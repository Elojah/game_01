package main

import (
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// Spawn moves the entity to spawn position and restore HP and MP.
func (svc *service) Spawn(id gulid.ID, e event.E) error {

	sp := e.Action.Spawn
	ts := e.ID.Time()

	// #Retrieve previous target state.
	target, err := svc.entity.Fetch(id, ts)
	if err != nil {
		return errors.Wrap(err, "spawn")
	}

	// #Check entity is dead
	if !target.Dead {
		return errors.Wrap(gerrors.ErrIsNotDead{EntityID: id.String()}, "spawn")
	}

	s, err := svc.entity.FetchSpawn(sp.ID)
	if err != nil {
		return errors.Wrap(err, "spawn")
	}

	// #Move target at spawn position
	target.Position = s.Position
	// #Restore HP & MP
	target.HP = target.MaxHP
	target.MP = target.MaxMP
	// #Target is not dead anymore \o/
	target.Dead = false

	// #Set entity new state.
	return errors.Wrap(svc.entity.Upsert(target, ts+1), "spawn")
}
