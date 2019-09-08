package ui

import (
	"sync"
	"time"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"

	"github.com/elojah/game_01/cmd/sandbox/reader"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/geometry"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

var _ ecs.System = (*ClientSystem)(nil)

type move struct {
	EntityID gulid.ID
	Position engo.Point
}

// Type returns move type.
func (move) Type() string {
	return "move"
}

// ClientSystem is an ecs system to handle move controls with player.
type ClientSystem struct {
	Reader   *reader.R
	Interval time.Duration

	entities sync.Map

	ticker *time.Ticker
}

func (s *ClientSystem) New(*ecs.World) {
	engo.Mailbox.Listen("move", func(message engo.Message) {
		input := message.(move)
		s.entities.Store(input.EntityID, input.Position)
	})
	s.ticker = time.NewTicker(s.Interval)
	go s.SendMove()
}

func (s *ClientSystem) SendMove() {
	for range s.ticker.C {
		s.entities.Range(func(key interface{}, value interface{}) bool {
			id := key.(gulid.ID)
			pos := value.(engo.Point)
			go s.Reader.Send(event.DTO{
				ID:    gulid.NewID(),
				Token: gulid.NewID(), // @TODO token system ?
				Query: event.Query{
					Move: &event.Move{
						Targets: []gulid.ID{id},
						Position: geometry.Position{
							SectorID: gulid.NewID(), //@TODO sectoe system ?
							Coord: geometry.Vec3{
								X: float64(pos.X),
								Y: float64(pos.Y),
							},
						},
					},
				},
			})
			s.entities.Delete(key)
			return true
		})
	}
}

func (s *ClientSystem) Remove(e ecs.BasicEntity) {}

func (s *ClientSystem) Update(dt float32) {}
