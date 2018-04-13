package main

import (
	"github.com/sirupsen/logrus"

	"github.com/elojah/game_01"
	"github.com/elojah/udp"
)

type actor struct {
	*logrus.Entry

	game.ActorService
}

// CreateActor .
func (a *actor) CreateActor(packet udp.Packet) error {
	return nil
}

// UpdateActor .
func (a *actor) UpdateActor(packet udp.Packet) error {
	return nil
}

// DeleteActor .
func (a *actor) DeleteActor(packet udp.Packet) error {
	return nil
}

// ListActor .
func (a *actor) ListActor(packet udp.Packet) error {
	return nil
}
