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
		return errors.Wrap(err, "retrieve entity")
	}

	// #Retrieve feedback.
	fb, err := a.FeedbackStore.GetFeedback(ft.ID)
	switch errors.Cause(err).(type) {
	case gerrors.ErrNotFound:
		return errors.Wrapf(err, "retrieve feedback")
	}

	// #Apply all ability components.
	if err := target.ApplyEffectFeedbacks(&ft.Source, fb.Effects); err != nil {
		return errors.Wrap(err, "apply feedback")
	}

	// #Set entity new state.
	return errors.Wrap(a.EntityStore.SetEntity(target, ts), "validate feeback")
}
