package main

import (
	"github.com/sirupsen/logrus"

	"github.com/elojah/game_01"
)

type actor struct {
	*logrus.Entry

	game.ActorService
}

// CreateActor .
func (a *actor) CreateActor(raw []byte) error {
	return nil
}

// UpdateActor .
func (a *actor) UpdateActor(raw []byte) error {
	return nil
}

// DeleteActor .
func (a *actor) DeleteActor(raw []byte) error {
	return nil
}

// ListActor .
func (a *actor) ListActor(raw []byte) error {
	return nil
}
