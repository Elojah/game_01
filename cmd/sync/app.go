package main

import (
	"github.com/elojah/game_01"
)

type app struct {
	game.QEventMapper
	game.TokenMapper
	game.EntityMapper

	tickRate uint32
}

func (a *app) Dial(c Config) error {
	a.tickRate = c.TickRate
	return nil
}
