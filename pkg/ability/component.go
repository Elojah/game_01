package ability

import (
	"github.com/elojah/game_01/pkg/entity"
)

// Affect a HealDirect on target.
func (c HealDirect) Affect(target *entity.E) ComponentFeedback {
	cf := ComponentFeedback{
		HealDirectFeedback: &HealDirectFeedback{
			Amount: c.Amount,
		},
	}
	target.HP += c.Amount
	return cf
}

// Affect a DamageDirect on target.
func (c DamageDirect) Affect(target *entity.E) ComponentFeedback {
	var cf ComponentFeedback
	if c.Amount >= target.HP {
		cf.SetValue(&DamageDirectFeedback{
			Amount: target.HP,
		})
		target.HP = 0
		return cf
	}
	cf.SetValue(&DamageDirectFeedback{
		Amount: c.Amount,
	})
	target.HP -= c.Amount
	return cf
}

// Affect a HealOverTime on target.
func (c HealOverTime) Affect(target *entity.E) ComponentFeedback {
	cf := ComponentFeedback{
		HealOverTimeFeedback: &HealOverTimeFeedback{
			Amount: c.Amount,
		},
	}
	return cf
}

// Affect a DamageOverTime on target.
func (c DamageOverTime) Affect(target *entity.E) ComponentFeedback {
	cf := ComponentFeedback{
		DamageOverTimeFeedback: &DamageOverTimeFeedback{
			Amount: c.Amount,
		},
	}
	return cf
}
