package ui

import (
	tl "github.com/JoelOtter/termloop"
)

type Term struct {
	g *tl.Game
}

func (t *Term) Close() error {
	return nil
}

// Dial initialize a Term.
func (t *Term) Dial(cfg Config) error {
	t.g = tl.NewGame()
	t.g.Screen().SetFps(cfg.FPS)

	l := tl.NewBaseLevel(tl.Cell{Bg: 76, Fg: 1})

	// lmap, err := ioutil.ReadFile(cfg.LevelFile)
	// if err != nil {
	// 	return err
	// }

	// parsers := make(map[string]tl.EntityParser)
	// parsers["Player"] = parsePlayer
	// err = tl.LoadLevelFromMap(string(lmap), parsers, l)

	t.g.Screen().SetLevel(l)
	t.g.Start()

	return nil
}
