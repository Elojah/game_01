package ability

import (
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/ulid"
)

// Affect applies ability a on target.
func (a A) Affect(target *entity.E) Feedback {
	fb := Feedback{
		ID:         ulid.NewID(),
		AbilityID:  a.ID,
		Components: make([]ComponentFeedback, len(a.Components)),
	}
	// for i, component := range a.Components {
	// fb.Components[i] = component.Affect(target)
	// }
	return fb
}

// Store is the communication interface for abilities.
type Store interface {
	SetAbility(A, ulid.ID) error
	GetAbility(Subset) (A, error)
	ListAbility(Subset) ([]A, error)
}

// Subset retrieves per EntityID+ID or list per EntityID.
type Subset struct {
	ID       ulid.ID
	EntityID ulid.ID
}
