package game

// AbilityComponent is a part of a skill.
type AbilityComponent interface {
	Affect(*Entity) AbilityFeedbackComponent
}

// HealDirect instant heals.
type HealDirect struct {
	Amount uint64
	Type   uint8
}

// Affect a HealDirect on target.
func (c HealDirect) Affect(target *Entity) AbilityFeedbackComponent {
	target.HP += c.Amount
	return HealDirectFeedback{}
}

// DamageDirect instant damage.
type DamageDirect struct {
	Amount uint64
	Type   uint8
}

// Affect a DamageDirect on target.
func (c DamageDirect) Affect(target *Entity) AbilityFeedbackComponent {
	if c.Amount >= target.HP {
		target.HP = 0
		return DamageDirectFeedback{}
	}
	target.HP -= c.Amount
	return DamageDirectFeedback{}
}

// HealOverTime heals per tick.
type HealOverTime struct {
	Amount    uint64
	Type      uint8
	Frequency uint64
	Duration  uint64
}

// Affect a HealOverTime on target.
func (c HealOverTime) Affect(target *Entity) AbilityFeedbackComponent {
	return HealOverTimeFeedback{}
}

// DamageOverTime inflicts damage per tick.
type DamageOverTime struct {
	Amount    uint64
	Type      uint8
	Frequency uint64
	Duration  uint64
}

// Affect a DamageOverTime on target.
func (c DamageOverTime) Affect(target *Entity) AbilityFeedbackComponent {
	return DamageOverTimeFeedback{}
}
