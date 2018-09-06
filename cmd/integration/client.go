package main

import (
	"github.com/elojah/game_01/pkg/geometry"
)

func expectClient(a *LogAnalyzer) (geometry.Position, error) {
	var pos geometry.Position
	err := a.Expect(func(s string) (bool, error) {
		return true, nil
	})
	return pos, err
}
