package main

import (
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/elojah/game_01/pkg/ability"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

func (a *app) PerformSource(id gulid.ID, e event.E) error {

	ps := e.Action.PerformSource
	ts := e.ID.Time()

	// #Retrieve source
	source, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrapf(err, "get entity %s", id.String())
	}

	// #Retrieve ability.
	ab, err := a.AbilityStore.GetAbility(source.ID, ps.AbilityID)
	switch errors.Cause(err).(type) {
	case gerrors.ErrNotFound:
		return errors.Wrapf(gerrors.ErrInsufficientACLs{
			Value:  -1,
			Source: source.ID.String(),
			Target: ps.AbilityID.String(),
		}, "get ability %s for %s", ps.AbilityID.String(), id.String())
	}
	if err != nil {
		return errors.Wrapf(err, "get ability %s for %s", ps.AbilityID.String(), id.String())
	}

	// #Check cast was not interrupted.
	if source.Cast == nil ||
		source.Cast.AbilityID != ps.AbilityID ||
		source.Cast.TS+ab.CastTime != ts {
		// normal behavior, don't return errors
		return nil
	}

	// #Set ability LastUsed
	// LastUsed is set in PerformSource and not in feedback because there are potentially X multiple targets.
	// If one or all target(s) fails (e.g: too far), we still apply to other targets.
	ab.LastUsed = ts
	if err := a.AbilityStore.SetAbility(ab, source.ID); err != nil {
		return errors.Wrapf(err, "set ability %s for %s", ab.ID.String(), source.ID.String())
	}

	// #For all ability components.
	var i uint64
	for cid := range ab.Components {

		// #Retrieve targets for this component.
		target, ok := ps.Targets[cid]
		if !ok {
			return errors.Wrapf(gerrors.ErrMissingTarget{
				AbilityID:   ab.ID.String(),
				ComponentID: cid,
			}, "ability %s component %s ", ab.ID.String(), cid)
		}

		// #Send event to all targets
		e := event.E{
			ID: gulid.NewTimeID(ts + i + 1),
			Action: event.Action{
				PerformTarget: &event.PerformTarget{
					AbilityID:   ab.ID,
					ComponentID: gulid.MustParse(cid),
					Source:      source,
				},
			},
			Trigger: e.ID,
		}

		if len(target.Positions) != 0 {
			return gerrors.ErrNotImplementedYet{Version: "0.2.0"}
		}

		var g errgroup.Group
		for _, id := range target.Entities {
			id := id
			g.Go(func() error {
				if err := a.EventQStore.PublishEvent(e, id); err != nil {
					return errors.Wrapf(err, "publish perform target event %s to target %s", e.ID.String(), target.String())
				}
				return nil
			})
		}
		if err := g.Wait(); err != nil {
			return err
		}
		i++
	}
	return nil
}

func (a *app) PerformTarget(id gulid.ID, e event.E) error {

	pt := e.Action.PerformTarget
	ts := e.ID.Time()

	// #Retrieve previous target state.
	target, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrapf(err, "get entity %s at max ts %d", id.String(), ts)
	}

	// #Retrieve ability.
	ab, err := a.AbilityStore.GetAbility(pt.Source.ID, pt.AbilityID)
	switch errors.Cause(err).(type) {
	case gerrors.ErrNotFound:
		return errors.Wrapf(gerrors.ErrInsufficientACLs{
			Value:  -1,
			Source: pt.Source.ID.String(),
			Target: pt.AbilityID.String(),
		}, "get ability %s for %s", pt.AbilityID.String(), pt.Source.ID.String())
	}
	if err != nil {
		return errors.Wrapf(err, "get ability %s for %s", pt.AbilityID.String(), pt.Source.ID.String())
	}

	// #Initialize feedback.
	fb := ability.Feedback{
		ID:          gulid.NewID(),
		AbilityID:   ab.ID,
		ComponentID: pt.ComponentID,
	}

	// #Check component validity.
	cid := pt.ComponentID.String()
	component, ok := ab.Components[cid]
	if !ok {
		return errors.Wrapf(gerrors.ErrMissingTarget{
			AbilityID:   ab.ID.String(),
			ComponentID: cid,
		}, "ability %s component %s", ab.ID.String(), cid)
	}

	// #Check distance between source and target
	dist, err := a.SectorService.Segment(pt.Source.Position, target.Position)
	if err != nil {
		return errors.Wrap(err, "consume target")
	}
	if dist > component.Range {
		return errors.Wrap(
			gerrors.ErrOutOfRange{
				Dist:  dist,
				Range: component.Range,
			},
			"consume target",
		)
	}

	// #Apply all ability components.
	if fb.Effects, err = target.ApplyEffects(&pt.Source, component.Effects); err != nil {
		return errors.Wrapf(err, "apply effects to target %s", target.ID.String())
	}

	// #Set entity new state.
	if err := a.EntityStore.SetEntity(target, ts); err != nil {
		return errors.Wrapf(err, "set entity %s", pt.Source.ID.String())
	}

	// #Set feedback.
	if err := a.FeedbackStore.SetFeedback(fb); err != nil {
		return errors.Wrapf(err, "set feedback %s from %s", fb.ID.String(), target.ID.String())
	}

	// #Publish feedback to source.
	return a.EventQStore.PublishEvent(event.E{
		ID: gulid.NewTimeID(ts + 1),
		Action: event.Action{
			PerformFeedback: &event.PerformFeedback{
				ID:     fb.ID,
				Target: target,
			},
		},
		Trigger: e.ID,
	}, pt.Source.ID)
}

func (a *app) PerformFeedback(id gulid.ID, e event.E) error {

	ft := e.Action.PerformFeedback
	ts := e.ID.Time()

	// #Retrieve previous source state.
	source, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrap(err, "retrieve entity")
	}

	// #Retrieve feedback.
	fb, err := a.FeedbackStore.GetFeedback(ft.ID)
	switch errors.Cause(err).(type) {
	case gerrors.ErrNotFound:
		return errors.Wrapf(err, "retrieve feedback")
	}

	// #Apply all ability components.
	if err := source.ApplyEffectFeedbacks(&ft.Target, fb.Effects); err != nil {
		return errors.Wrap(err, "apply feedback")
	}

	// #Set entity new state.
	return errors.Wrap(a.EntityStore.SetEntity(source, ts), "validate feeback")
}
