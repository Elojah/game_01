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

func (c *ControlSystem) Add(e *Entity) {
	c.entity = e
}

func (c *ControlSystem) Remove(e ecs.BasicEntity) {
	if c.entity != nil && e.ID() == c.entity.ID() {
		c.entity = nil
	}
}

func (c *ControlSystem) Update(dt float32) {
	for _, input := range []struct {
		name string
		f    func()
	}{
		{
			name: "walk_up",
			f:    func() { c.entity.SpaceComponent.Position.Y -= 5 },
		},

		{
			name: "walk_down",
			f:    func() { c.entity.SpaceComponent.Position.Y += 5 },
		},

		{
			name: "walk_left",
			f:    func() { c.entity.SpaceComponent.Position.X -= 5 },
		},

		{
			name: "walk_right",
			f:    func() { c.entity.SpaceComponent.Position.X += 5 },
		},
	} {
		if engo.Input.Button(input.name).Down() {
			if c.entity.AnimationComponent.CurrentAnimation.Name != input.name {
				c.entity.AnimationComponent.SelectAnimationByName(input.name)
			}
			input.f()
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
