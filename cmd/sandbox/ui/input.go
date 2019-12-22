package ui

import "time"

type Input string

const (
	refresh = 8 * time.Millisecond

	WalkUp    Input = "walk_up"
	WalkDown  Input = "walk_down"
	WalkLeft  Input = "walk_left"
	WalkRight Input = "walk_right"
)

// InputChan represent a throttled channel of input reading each 8ms.
type InputChan struct {
	ticker <-chan time.Time
}

// NewInputChan returns a valid InputChan.
func NewInputChan() InputChan {
	return InputChan{
		ticker: time.Tick(refresh),
	}
}

func (c InputChan) Run(f func()) {
	for _ = range c.ticker {
		f()
	}
}

// InputAction map input and entity action.
type InputAction struct {
	Input  Input
	Action func(*Entity)
}

var (
	actions = []InputAction{
		{
			Input:  WalkUp,
			Action: func(e *Entity) { e.SpaceComponent.Position.Y -= 5 },
		},
		{
			Input:  WalkDown,
			Action: func(e *Entity) { e.SpaceComponent.Position.Y += 5 },
		},
		{
			Input:  WalkLeft,
			Action: func(e *Entity) { e.SpaceComponent.Position.X -= 5 },
		},
		{
			Input:  WalkRight,
			Action: func(e *Entity) { e.SpaceComponent.Position.X += 5 },
		},
	}
)
