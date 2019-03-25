package main

import (
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

func (a *app) CastSource(id ulid.ID, e event.E) error {

	cs := e.Action.CastSource
	ts := e.ID.Time()

	// #Check permission source/token
	if err := a.EntityPermissionService.CheckPermission(e.Token, id); err != nil {
		return errors.Wrap(err, "cast source")
	}

	// #Retrieve ability.
	ab, err := a.AbilityStore.GetAbility(id, cs.AbilityID)
	switch errors.Cause(err).(type) {
	case gerrors.ErrNotFound:
		return errors.Wrap(
			gerrors.ErrInsufficientACLs{
				Value:  -1,
				Source: id.String(),
				Target: cs.AbilityID.String(),
			},
			"cast source",
		)
	}
	if err != nil {
		return errors.Wrap(err, "cast source")
	}

	// #Retrieve entity
	source, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrap(err, "cast source")
	}

	// #Check if entity is alive
	if source.Dead {
		return errors.Wrap(gerrors.ErrIsDead{EntityID: id.String()}, "cast source")
	}

	// #Check MP consumption
	if source.MP < ab.MPConsumption {
		return errors.Wrap(
			gerrors.ErrMissingMP{
				EntityID:      id.String(),
				MPLeft:        source.MP,
				AbilityID:     ab.ID.String(),
				MPConsumption: ab.MPConsumption,
			}, "cast source",
		)
	}

	// #Check CD validity. if LastUsed + CD < now.
	if ts-ab.CD < ab.LastUsed {
		return errors.Wrap(
			gerrors.ErrAbilityCDDown{
				EntityID:  source.ID.String(),
				AbilityID: ab.ID.String(),
				TS:        ts,
				LastUsed:  ab.LastUsed,
				CD:        ab.CD,
			},
			"cast source",
		)
	}

	// Check targets validity (number and position number).
	for cid, component := range ab.Components {
		targets, ok := cs.Targets[cid]
		if !ok {
			return errors.Wrap(gerrors.ErrMissingTarget{
				AbilityID:   ab.ID.String(),
				ComponentID: cid,
			}, "cast source")
		}
		if uint64(len(targets.Entities)) > component.NTargets {
			return errors.Wrap(gerrors.ErrTooManyTargets{
				NTargets:    len(targets.Entities),
				Max:         component.NTargets,
				AbilityID:   ab.ID.String(),
				ComponentID: cid,
			}, "cast source")
		}
		if uint64(len(targets.Positions)) > component.NPositions {
			return errors.Wrap(gerrors.ErrTooManyTargets{
				NTargets:    len(targets.Positions),
				Max:         component.NPositions,
				AbilityID:   ab.ID.String(),
				ComponentID: cid,
			}, "cast source")
		}
	}

	// #Set entity new state with decreased MP and casting up.
	source.CastAbility(ab, ts)
	if err := a.EntityStore.SetEntity(source, ts); err != nil {
		return errors.Wrap(err, "cast source")
	}

	// #Publish casted event to event set.
	return errors.Wrap(a.EventQStore.PublishEvent(
		event.E{
			ID: ulid.NewTimeID(ts + ab.CastTime),
			Action: event.Action{
				PerformSource: &event.PerformSource{
					AbilityID: cs.AbilityID,
					Targets:   cs.Targets,
				},
			},
			Trigger: e.ID,
		}, id),
		"cast source",
	)
}
