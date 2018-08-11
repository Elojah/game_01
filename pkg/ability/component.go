package ability

import (
	"github.com/elojah/game_01/pkg/entity"
)

// Affect a HealDirect on target.
func (c HealDirect) Affect(target *entity.E) FeedbackComponent {
	target.HP += c.Amount
	return HealDirectFeedback{}
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

// Affect a HealOverTime on target.
func (c HealOverTime) Affect(target *entity.E) FeedbackComponent {
	return HealOverTimeFeedback{}
}

// Affect a DamageOverTime on target.
func (c DamageOverTime) Affect(target *entity.E) FeedbackComponent {
	return DamageOverTimeFeedback{}
}
