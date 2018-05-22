package game

// SkillType represents the skill type.
type SkillType = ID

// Skill represents a skill.
type Skill struct {
	ID            ID        `json:"id"`
	Type          SkillType `json:"type"`
	Name          string    `json:"name"`
	MPConsumption uint64    `json:"mp_consumption"`
	DirectDamage  uint64    `json:"direct_damage"`
	DirectHeal    uint64    `json:"direct_heal"`
	CD            uint32    `json:"cd"`
	CurrentCD     uint32    `json:"current_cd"`
}

// SkillMapper is the communication interface for skills.
type SkillMapper interface {
	SetSkill(Skill, ID) error
	GetSkill(SkillSubset) (Skill, error)
	ListSkill(SkillSubset) ([]Skill, error)
}

// SkillSubset retrieves per EntityID+ID or list per EntityID.
type SkillSubset struct {
	ID       ID
	EntityID ID
}

// SkillEffect represents the effects a skill had on a target.
type SkillEffect struct {
	DamageDone float64
	HealDone   float64
}

// Apply applies skill s on target.
func (s Skill) Apply(target Entity) error {
	return nil
}
