package ability

import (
	"github.com/elojah/game_01/pkg/entity"
)

// Component is a part of a skill.
type Component interface {
	Affect(*entity.E) FeedbackComponent
}

// HealDirect instant heals.
type HealDirect struct {
	Amount uint64
	Type   uint8
}

// Affect a HealDirect on target.
func (c HealDirect) Affect(target *entity.E) FeedbackComponent {
	target.HP += c.Amount
	return HealDirectFeedback{}
}

// DamageDirect instant damage.
type DamageDirect struct {
	Amount uint64
	Type   uint8
}

// Affect a DamageDirect on target.
func (c DamageDirect) Affect(target *entity.E) FeedbackComponent {
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
func (c HealOverTime) Affect(target *entity.E) FeedbackComponent {
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
func (c DamageOverTime) Affect(target *entity.E) FeedbackComponent {
	return DamageOverTimeFeedback{}
}
