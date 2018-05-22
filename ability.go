package game

// AbilityType represents the ability type.
type AbilityType = ID

// Ability represents a ability.
type Ability struct {
	ID            ID          `json:"id"`
	Type          AbilityType `json:"type"`
	Name          string      `json:"name"`
	MPConsumption uint64      `json:"mp_consumption"`
	DirectDamage  uint64      `json:"direct_damage"`
	DirectHeal    uint64      `json:"direct_heal"`
	CD            uint32      `json:"cd"`
	CurrentCD     uint32      `json:"current_cd"`
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

// AbilityFeedback represents the effects a ability had on a target.
type AbilityFeedback struct {
}

// Apply applies ability s on target.
func (s Ability) Apply(target Entity) error {
	return nil
}
