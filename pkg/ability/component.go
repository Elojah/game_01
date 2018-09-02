package ability

import (
	gerrors "github.com/elojah/game_01/pkg/errors"
)

// CheckTargets check if targets number is valid. Cast check only, no range check (done at perform).
func (c Component) CheckTargets(targets Targets) error {
	if uint64(len(targets.Entities)) > c.NTargets {
		return gerrors.ErrTooManyTargets
	}
	if uint64(len(targets.Positions)) > c.NPositions {
		return gerrors.ErrTooManyTargets
	}
	return nil
}
