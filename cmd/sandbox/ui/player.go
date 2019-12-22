package ui

import (
	tl "github.com/JoelOtter/termloop"
)

type Player struct {
	*tl.Entity
}

func (p *Player) Tick(ev tl.Event) {
	// Enable arrow key movement
	if ev.Type == tl.EventKey {
		x, y := p.Position()
		switch ev.Key {
		case tl.KeyArrowRight:
			x++
		case tl.KeyArrowLeft:
			x--
		case tl.KeyArrowUp:
			y--
		case tl.KeyArrowDown:
			y++
		}
		p.SetPosition(x, y)
	}
}

func UnmarshalPlayer(data map[string]interface{}) tl.Drawable {
	e := tl.NewEntity(
		int(data["x"].(float64)),
		int(data["y"].(float64)),
		1, 1,
	)
	e.SetCell(0, 0, &tl.Cell{
		Ch: 'o',
		Fg: tl.Attr(42),
	})
	return &Player{e}
}
