package game

// AbilityComponent is a part of a skill.
type AbilityComponent interface {
	Apply(*Entity) AbilityFeedback
}

// HealDirect instant heals.
type HealDirect struct {
	Amount uint64
	Type   uint8
}

// Apply a HealDirect on target.
func (c HealDirect) Apply(target *Entity) AbilityFeedback {
	target.HP += c.Amount
	return AbilityFeedback{}
}

// DamageDirect instant damage.
type DamageDirect struct {
	Amount uint64
	Type   uint8
}

// Apply a DamageDirect on target.
func (c DamageDirect) Apply(target *Entity) AbilityFeedback {
	if c.Amount >= target.HP {
		target.HP = 0
		return AbilityFeedback{}
	}
	target.HP -= c.Amount
	return AbilityFeedback{}
}

// HealOverTime heals per tick.
type HealOverTime struct {
	Amount    uint64
	Type      uint8
	Frequency uint64
	Duration  uint64
}

// Apply a HealOverTime on target.
func (c HealOverTime) Apply(target *Entity) AbilityFeedback {
	return AbilityFeedback{}
}

// DamageOverTime inflicts damage per tick.
type DamageOverTime struct {
	Amount    uint64
	Type      uint8
	Frequency uint64
	Duration  uint64
}

// Apply a DamageOverTime on target.
func (c DamageOverTime) Apply(target *Entity) AbilityFeedback {
	return AbilityFeedback{}
}
