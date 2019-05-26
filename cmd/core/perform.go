package main

import (
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/elojah/game_01/pkg/ability"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

func (svc *service) PerformSource(id gulid.ID, e event.E) error {

	ps := e.Action.PerformSource
	ts := e.ID.Time()

	// #Retrieve source
	source, err := svc.entity.Fetch(id, ts)
	if err != nil {
		return errors.Wrap(err, "perform source")
	}

	// #Check if entity is alive
	if source.Dead {
		return errors.Wrap(gerrors.ErrIsDead{EntityID: id.String()}, "perform source")
	}

	// #Retrieve ability.
	ab, err := svc.ability.Fetch(source.ID, ps.AbilityID)
	if err != nil {
		return errors.Wrap(err, "perform source")
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
	if err := svc.ability.Upsert(ab, source.ID); err != nil {
		return errors.Wrap(err, "perform source")
	}

	// #For all ability components.
	var i uint64
	for cid := range ab.Components {

		// #Retrieve targets for this component.
		target, ok := ps.Targets[cid]
		if !ok {
			return errors.Wrap(
				gerrors.ErrMissingTarget{
					AbilityID:   ab.ID.String(),
					ComponentID: cid,
				},
				"perform source",
			)
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
			return errors.Wrap(gerrors.ErrNotImplementedYet{Version: "0.2.0"}, "perform source")
		}

		var g errgroup.Group
		for _, id := range target.Entities {
			id := id
			g.Go(func() error {
				if err := svc.event.Publish(e, id); err != nil {
					return errors.Wrap(err, "perform source")
				}
				return nil
			})
		}
		if err := g.Wait(); err != nil {
			return err
		}
		i++
	}

	// #Remove cast state from source
	source.Cast = nil
	return errors.Wrap(svc.entity.Upsert(source, ts+1), "perform source")
}

func (svc *service) PerformTarget(id gulid.ID, e event.E) error {

	pt := e.Action.PerformTarget
	ts := e.ID.Time()

	// #Retrieve previous target state.
	target, err := svc.entity.Fetch(id, ts)
	if err != nil {
		return errors.Wrapf(err, "perform target")
	}

	// #Retrieve ability.
	ab, err := svc.ability.Fetch(pt.Source.ID, pt.AbilityID)
	if err != nil {
		return errors.Wrapf(err, "perform target")
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
		return errors.Wrap(
			gerrors.ErrMissingTarget{
				AbilityID:   ab.ID.String(),
				ComponentID: cid,
			},
			"perform target",
		)
	}

	// #Check distance between source and target
	dist, err := svc.sector.Segment(pt.Source.Position, target.Position)
	if err != nil {
		return errors.Wrap(err, "perform target")
	}
	if dist > component.Range {
		return errors.Wrap(
			gerrors.ErrOutOfRange{
				Dist:  dist,
				Range: component.Range,
			},
			"perform target",
		)
	}

	wasDead := target.Dead

	// #Apply all ability components.
	if fb.Effects, err = target.ApplyEffects(&pt.Source, component.Effects); err != nil {
		return errors.Wrapf(err, "perform target")
	}

	// #Add svc spawn event to svc freshly dead entity.
	// wasDead is svc guard against multiple spawn events.
	if !wasDead && target.Dead {
		sp, err := svc.entity.FetchSpawn(target.SpawnID)
		if err != nil {
			return errors.Wrap(err, "perform target")
		}
		if err := svc.event.Publish(event.E{
			ID: gulid.NewTimeID(ts + sp.Duration + 1),
			Action: event.Action{
				Spawn: &event.Spawn{
					ID: target.SpawnID,
				},
			},
			Trigger: e.ID,
		}, id); err != nil {
			return errors.Wrapf(err, "perform target")
		}
	}

	// #Set entity new state.
	if err := svc.entity.Upsert(target, ts+1); err != nil {
		return errors.Wrapf(err, "perform target")
	}

	// #Set feedback.
	if err := svc.ability.UpsertFeedback(fb); err != nil {
		return errors.Wrapf(err, "perform target")
	}

	// #Publish feedback to source.
	return errors.Wrap(svc.event.Publish(event.E{
		ID: gulid.NewTimeID(ts + 1),
		Action: event.Action{
			PerformFeedback: &event.PerformFeedback{
				ID:     fb.ID,
				Target: target,
			},
		},
		Trigger: e.ID,
	}, pt.Source.ID),
		"perform target",
	)
}

func (svc *service) PerformFeedback(id gulid.ID, e event.E) error {

	pf := e.Action.PerformFeedback
	ts := e.ID.Time()

	// #Retrieve previous source state.
	source, err := svc.entity.Fetch(id, ts)
	if err != nil {
		return errors.Wrap(err, "perform feedback")
	}

	// #Retrieve feedback.
	fb, err := svc.ability.FetchFeedback(pf.ID)
	switch errors.Cause(err).(type) {
	case gerrors.ErrNotFound:
		return errors.Wrap(err, "perform feedback")
	}

	// #Apply all ability components.
	if err := source.ApplyEffectFeedbacks(&pf.Target, fb.Effects); err != nil {
		return errors.Wrap(err, "perform feedback")
	}

	// #Set entity new state.
	return errors.Wrap(svc.entity.Upsert(source, ts+1), "perform feedback")
}
