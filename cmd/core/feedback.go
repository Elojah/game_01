package main

import (
	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
)

func (a *app) FeedbackTarget(id ulid.ID, e event.E) error {

	ft := e.Action.FeedbackTarget
	ts := e.ID.Time()

	// #Retrieve previous target state.
	target, err := a.EntityStore.GetEntity(id, ts)
	if err != nil {
		return errors.Wrapf(err, "get entity %s at max ts %d", id.String(), ts)
	}

	// #Retrieve feedback.
	fb, err := a.FeedbackStore.GetFeedback(ft.ID)
	if err == gerrors.ErrNotFound {
		return errors.Wrapf(gerrors.ErrNotFound, "get feedback %s set by %s", ft.ID.String(), ft.Source.ID.String())
	}

	// #Apply all ability components.
	if err := target.ApplyEffectFeedbacks(&ft.Source, fb.Effects); err != nil {
		return errors.Wrapf(err, "apply effects to target %s", target.ID.String())
	}

	// #Set entity new state.
	return errors.Wrapf(a.EntityStore.SetEntity(target, ts), "set entity %s", ft.Source.ID.String())
}
