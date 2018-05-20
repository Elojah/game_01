package game

// SkillType represents the skill type.
type SkillType = ID

// Skill represents a skill.
type Skill struct {
	ID            ID
	Type          SkillType
	Name          string
	MPConsumption uint64
	DirectDamage  uint64
	DirectHeal    uint64
	CD            uint32
	CurrentCD     uint32
}

// SkillService is the communication interface for skills.
type SkillService interface {
	SetSkill(Skill, ID) error
	GetSkill(SkillSubset) (Skill, error)
	ListSkill(SkillSubset) ([]Skill, error)
}

// SkillSubset retrieves per EntityID+ID or list per EntityID.
type SkillSubset struct {
	ID       ID
	EntityID ID
}
