package ui

import (
	"io/ioutil"

	tl "github.com/JoelOtter/termloop"
)

type Term struct {
	g *tl.Game

	Player *Player
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

	if err := tl.LoadLevelFromMap(string(lmap), nil, l); err != nil {
		return err
	}

	t.Player = NewPlayer()
	t.Player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '옷'})
	l.AddEntity(t.Player)

	t.g.Screen().SetLevel(l)
	go t.g.Start()

	return nil
}
