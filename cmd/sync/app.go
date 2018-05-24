package main

import (
	"github.com/elojah/game_01"
)

type app struct {
	game.QEventMapper
	game.TokenMapper
	game.EntityMapper
}

func (a *app) Dial(c Config) error {
	return nil
}
