package ability

import (
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/ulid"
)

// Type represents the ability type.
type Type = ulid.ID

// A A represents a ability.
type A struct {
	ID   ulid.ID `json:"id"`
	Type Type    `json:"type"`
	Name string  `json:"name"`

	MPConsumption uint64 `json:"mp_consumption"`
	CD            uint32 `json:"cd"`
	CurrentCD     uint32 `json:"current_cd"`

	Components []Component `json:"components"`
}

// Affect applies ability a on target.
func (a A) Affect(target *entity.E) Feedback {
	var fb Feedback
	fb.AbilityID = a.ID
	fb.Components = make([]FeedbackComponent, len(a.Components))
	for i, component := range a.Components {
		fb.Components[i] = component.Affect(target)
	}
	return fb
}

// Mapper is the communication interface for abilities.
type Mapper interface {
	SetAbility(A, ulid.ID) error
	GetAbility(Subset) (A, error)
	ListAbility(Subset) ([]A, error)
}

// Subset retrieves per EntityID+ID or list per EntityID.
type Subset struct {
	ID       ulid.ID
	EntityID ulid.ID
}
