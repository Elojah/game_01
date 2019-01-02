package main

import (
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/elojah/game_01/pkg/account"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

func (a *app) MoveSource(id ulid.ID, e event.E) error {

	ms := e.Action.MoveSource
	ts := e.ID.Time()

	// #Check permission token/source.
	permission, err := a.EntityPermissionStore.GetPermission(e.Token.String(), id.String())
	if err == gerrors.ErrNotFound || (err != nil && account.ACL(permission.Value) != account.Owner) {
		return errors.Wrapf(gerrors.ErrInsufficientACLs, "get permission token %s for %s", e.Token.String(), id.String())
	}
	if err != nil {
		return errors.Wrapf(err, "get permission token %s for %s", e.Token.String(), id.String())
	}

	// #TODO Check if source is not stun or forbidden to move other entities

	// #For all targets.
	var g errgroup.Group
	for _, target := range ms.TargetIDs {
		target := target
		g.Go(func() error {
			// #Check permission source/target.
			permission, err := a.EntityPermissionStore.GetPermission(id.String(), target.String())
			if err == gerrors.ErrNotFound || (err != nil && account.ACL(permission.Value) != account.Owner) {
				return errors.Wrapf(gerrors.ErrInsufficientACLs, "get permission token %s for %s", id.String(), target.String())
			}
			if err != nil {
				return errors.Wrapf(err, "get permission token %s for %s", id.String(), target.String())
			}

			// #Publish move event to target.
			e := event.E{
				ID: ulid.NewTimeID(ts + 1),
				Action: event.Action{
					MoveTarget: &event.MoveTarget{
						SourceID: id,
						Position: ms.Position,
					},
				},
			}
			if err := a.EventQStore.PublishEvent(e, target); err != nil {
				return errors.Wrapf(err, "publish move target event %s to target %s", e.ID.String(), target.String())
			}
			return nil
		})
	}

	return g.Wait()
}

func (a *app) MoveTarget(id ulid.ID, e event.E) error {

	mt := e.Action.MoveTarget
	ts := e.ID.Time()

	// #Retrieve previous state target.
	target, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrapf(err, "get entity %s at max ts %d", id.String(), ts)
	}

	if target, err = a.SectorService.Move(target, mt.Position); err != nil {
		return errors.Wrapf(err, "move target %s", id.String())
	}

	return errors.Wrapf(a.EntityStore.SetEntity(target, ts), "set entity %s for ts %d", target.ID.String(), ts)
}
