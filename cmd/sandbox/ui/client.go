package ui

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"

	"github.com/elojah/game_01/cmd/sandbox/reader"
	"github.com/elojah/game_01/pkg/event"
)

var _ ecs.System = (*ClientSystem)(nil)

type input struct {
	event.DTO
}

// Type returns input type.
func (input) Type() string {
	return "input"
}

// ClientSystem is an ecs system to handle input controls with player.
type ClientSystem struct {
	Reader *reader.R
}

func (s *ClientSystem) New(*ecs.World) {
	engo.Mailbox.Listen("input", func(message engo.Message) {
		in := message.(input)
		go s.Reader.Send(in.DTO)
	})
}

func (s *ClientSystem) Remove(e ecs.BasicEntity) {}

func (s *ClientSystem) Update(dt float32) {}
