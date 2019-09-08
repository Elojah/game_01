package ui

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
)

var _ ecs.System = (*ControlSystem)(nil)

// ControlSystem is an ecs system to handle input controls with player.
type ControlSystem struct {
	entity *Entity
}

func (s *ControlSystem) Add(e *Entity) {
	s.entity = e
}

func (s *ControlSystem) Remove(e ecs.BasicEntity) {
	if s.entity != nil && e.ID() == s.entity.ID() {
		s.entity = nil
	}
}

func (s *ControlSystem) Update(dt float32) {
	for _, in := range []struct {
		name string
		f    func(*Entity)
	}{
		{
			name: "walk_up",
			f:    func(e *Entity) { e.SpaceComponent.Position.Y -= 5 },
		},
		{
			name: "walk_down",
			f:    func(e *Entity) { e.SpaceComponent.Position.Y += 5 },
		},
		{
			name: "walk_left",
			f:    func(e *Entity) { e.SpaceComponent.Position.X -= 5 },
		},
		{
			name: "walk_right",
			f:    func(e *Entity) { e.SpaceComponent.Position.X += 5 },
		},
	} {
		if engo.Input.Button(in.name).JustPressed() {
			if s.entity.AnimationComponent.CurrentAnimation.Name != in.name {
				s.entity.AnimationComponent.SelectAnimationByName(in.name)
			}
		}
		if engo.Input.Button(in.name).Down() {
			in.f(s.entity)
			// Dispatch event to client
			engo.Mailbox.Dispatch(move{
				EntityID: s.entity.GID,
				Position: s.entity.SpaceComponent.Position,
			})
		}
	}
}

// SetupControls setup default controls.
func SetupControls() {
	engo.Input.RegisterButton("walk_up", engo.KeyW)
	engo.Input.RegisterButton("walk_left", engo.KeyA)
	engo.Input.RegisterButton("walk_down", engo.KeyS)
	engo.Input.RegisterButton("walk_right", engo.KeyD)
}
