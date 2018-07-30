package ability

import (
	"encoding/json"

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
	fb := Feedback{
		ID:         ulid.NewID(),
		AbilityID:  a.ID,
		Components: make([]FeedbackComponent, len(a.Components)),
	}
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

// UnmarshalJSON adds a new field for components `struct` to determine component type.
func (a *Template) UnmarshalJSON(raw []byte) error {
	type aliasA A
	var alias struct {
		aliasA
		Components []json.RawMessage `json:"components"`
	}
	if err := json.Unmarshal(raw, &alias); err != nil {
		return err
	}

	*a = Template(alias.aliasA)
	a.Components = make([]Component, len(alias.Components))

	for i, component := range alias.Components {
		var componentStruct struct {
			Structure string `json:"struct"`
		}
		if err := json.Unmarshal([]byte(component), &componentStruct); err != nil {
			return err
		}
		switch componentStruct.Structure {
		case "HealDirect":
			var s HealDirect
			if err := json.Unmarshal([]byte(component), &s); err != nil {
				return err
			}
			a.Components[i] = s
		case "DamageDirect":
			var s DamageDirect
			if err := json.Unmarshal([]byte(component), &s); err != nil {
				return err
			}
			a.Components[i] = s
		case "HealOverTime":
			var s HealOverTime
			if err := json.Unmarshal([]byte(component), &s); err != nil {
				return err
			}
			a.Components[i] = s
		case "DamageOverTime":
			var s DamageOverTime
			if err := json.Unmarshal([]byte(component), &s); err != nil {
				return err
			}
			a.Components[i] = s
		default:
			return &json.UnsupportedValueError{Str: "struct"}
		}
	}
	return nil
}
