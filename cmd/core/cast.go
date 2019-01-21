package main

import (
	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/account"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

func (a *app) CastSource(id ulid.ID, e event.E) error {

	cs := e.Action.CastSource
	ts := e.ID.Time()

	// #Check permission token/source.
	permission, err := a.EntityPermissionStore.GetPermission(e.Token.String(), id.String())
	if errors.Cause(err).(type) == gerrors.ErrNotFound || (err == nil && account.ACL(permission.Value) != account.Owner) {
		return errors.Wrapf(gerrors.ErrInsufficientACLs{
			Value:  permission.Value,
			Source: e.Token.String(),
			Target: id.String(),
		}, "get permission token %s for %s", e.Token.String(), id.String())
	}
	if err != nil {
		return errors.Wrapf(err, "get permission token %s for %s", e.Token.String(), id.String())
	}

	// #Retrieve ability.
	ab, err := a.AbilityStore.GetAbility(id, cs.AbilityID)
	if errors.Cause(err).(type) == gerrors.ErrNotFound {
		return errors.Wrapf(gerrors.ErrInsufficientACLs{
			Value:  -1,
			Source: id.String(),
			Target: cs.AbilityID.String(),
		}, "get ability %s for %s", cs.AbilityID.String(), id.String())
	}
	if err != nil {
		return errors.Wrapf(err, "get ability %s for %s", cs.AbilityID.String(), id.String())
	}

	// #Check MP consumption
	source, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrapf(err, "get entity %s at max ts %d", id.String(), ts)
	}
	if source.MP < ab.MPConsumption {
		return errors.Wrapf(
			gerrors.ErrInvalidAction{Action: "cast_source"},
			"entity %s with MP %d for ability %s with MP: %d",
			id.String(),
			source.MP,
			ab.ID.String(),
			ab.MPConsumption,
		)
	}

	// #Check CD validity. if LastUsed + CD < now.
	if ts-ab.CD < ab.LastUsed {
		return errors.Wrapf(gerrors.ErrInvalidAction{Action: "cast_source"}, "cd down for skill %s ", ab.ID.String())
	}

	// Check targets validity (number and position number).
	for cid, component := range ab.Components {
		targets, ok := cs.Targets[cid]
		if !ok {
			return errors.Wrapf(gerrors.ErrMissingTarget{
				AbilityID:   ab.ID.String(),
				ComponentID: cid.String(),
			}, "ability %s component %s", ab.ID.String(), cid)
		}
		if uint64(len(targets.Entities)) > component.NTargets {
			return errors.Wrapf(gerrors.ErrTooManyTargets{
				NTargets:    len(targets.Entities),
				Max:         component.NTargets,
				AbilityID:   ab.ID.String(),
				ComponentID: cid,
			}, "%d entities for %d max for ability %s component %s", len(targets.Entities), component.NTargets, ab.ID.String(), cid)
		}
		if uint64(len(targets.Positions)) > component.NPositions {
			return errors.Wrapf(gerrors.ErrTooManyTargets{
				NTargets:    len(targets.Entities),
				Max:         component.NTargets,
				AbilityID:   ab.ID.String(),
				ComponentID: cid,
			}, "%d positions for %d max for ability %s component %s", len(targets.Positions), component.NPositions, ab.ID.String(), cid)
		}
	}

	// #Set entity new state with decreased MP and casting up.
	source.CastAbility(ab, ts)
	if err := a.EntityStore.SetEntity(source, ts); err != nil {
		return errors.Wrapf(err, "set entity %s", source.ID.String())
	}

	// #Publish casted event to event set.
	e = event.E{
		ID: ulid.NewTimeID(ts + ab.CastTime),
		Action: event.Action{
			PerformSource: &event.PerformSource{
				AbilityID: cs.AbilityID,
				Targets:   cs.Targets,
			},
		},
		Trigger: e.ID,
	}
	if err := a.EventQStore.PublishEvent(e, id); err != nil {
		return errors.Wrapf(err, "set casted event %s", e.ID.String())
	}

	return nil
}
