package main

import (
	"fmt"

	"github.com/elojah/game_01/pkg/geometry"
)

func expectClient(a *LogAnalyzer) (geometry.Position, error) {
	var pos geometry.Position
	err := a.Expect(func(s string) (bool, error) {
		fmt.Println(s)
		return true, nil
	})
	return pos, err
}
