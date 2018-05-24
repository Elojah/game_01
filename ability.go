package game

// AbilityType represents the ability type.
type AbilityType = ID

// Ability represents a ability.
type Ability struct {
	ID   ID          `json:"id"`
	Type AbilityType `json:"type"`
	Name string      `json:"name"`

	MPConsumption uint64 `json:"mp_consumption"`
	CD            uint32 `json:"cd"`
	CurrentCD     uint32 `json:"current_cd"`

	Components []AbilityComponent `json:"components"`
}

// Affect applies ability a on target.
func (a Ability) Affect(target *Entity) AbilityFeedback {
	var afb AbilityFeedback
	afb.AbilityID = a.ID
	afb.Components = make([]AbilityFeedbackComponent, len(a.Components))
	for i, component := range a.Components {
		afb.Components[i] = component.Affect(target)
	}
	return afb
}

// AbilityMapper is the communication interface for abilitys.
type AbilityMapper interface {
	SetAbility(Ability, ID) error
	GetAbility(AbilitySubset) (Ability, error)
	ListAbility(AbilitySubset) ([]Ability, error)
}

// AbilitySubset retrieves per EntityID+ID or list per EntityID.
type AbilitySubset struct {
	ID       ID
	EntityID ID
}
