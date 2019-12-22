package ui

import (
	tl "github.com/JoelOtter/termloop"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/geometry"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

type Player struct {
	id gulid.ID
	*tl.Entity
}

func NewPlayer() *Player {
	return &Player{
		id:     gulid.NewID(),
		Entity: tl.NewEntity(1, 1, 1, 1),
	}
}

func (p *Player) ID() gulid.ID {
	return p.id
}

func (p *Player) Query() event.Query {
	x, y := p.Position()
	return event.Query{
		Move: &event.Move{
			Targets: []gulid.ID{p.id},
			Position: geometry.Position{
				SectorID: gulid.NewID(), //TODO
				Coord: geometry.Vec3{
					X: float64(x),
					Y: float64(y),
					Z: 0,
				},
			},
		},
	}
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
