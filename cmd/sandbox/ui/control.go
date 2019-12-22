package ui

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
)

var _ ecs.System = (*ControlSystem)(nil)

// ControlSystem is an ecs system to handle input controls with player.
type ControlSystem struct {
	entity *Entity
	move   InputChan
}

// SetupControls setup default controls.
func (s *ControlSystem) Setup() {
	engo.Input.RegisterButton(string(WalkUp), engo.KeyW)
	engo.Input.RegisterButton(string(WalkLeft), engo.KeyA)
	engo.Input.RegisterButton(string(WalkDown), engo.KeyS)
	engo.Input.RegisterButton(string(WalkRight), engo.KeyD)

	s.move = NewInputChan()
	go s.move.Run(s.HandleInput)
}

func (s *ControlSystem) Add(e *Entity) {
	s.entity = e
}

func (s *ControlSystem) Remove(e ecs.BasicEntity) {
	if s.entity != nil && e.ID() == s.entity.ID() {
		s.entity = nil
	}
}

func (s *ControlSystem) Update(dt float32) {}

func (s *ControlSystem) HandleInput() {
	for _, ai := range actions {
		if engo.Input.Button(string(ai.Input)).Down() {
			ai.Action(s.entity)
			// Dispatch event to client
			engo.Mailbox.Dispatch(move{
				EntityID: s.entity.GID,
				Position: s.entity.SpaceComponent.Position,
			})
		}
	}
}
