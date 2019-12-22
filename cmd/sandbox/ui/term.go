package ui

import (
	"io/ioutil"

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

	l := tl.NewBaseLevel(tl.Cell{Bg: 0, Fg: 220})

	lmap, err := ioutil.ReadFile(cfg.LevelFile)
	if err != nil {
		return err
	}

	parsers := make(map[string]tl.EntityParser)
	parsers["player"] = UnmarshalPlayer
	if err := tl.LoadLevelFromMap(string(lmap), parsers, l); err != nil {
		return err
	}

	t.g.Screen().SetLevel(l)
	t.g.Start()

	return nil
}
