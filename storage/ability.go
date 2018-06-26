package storage

import (
	game "github.com/elojah/game_01"
	"github.com/elojah/game_01/pkg/ability"
)

// NewHealDirect converts a domain HealDirect into a storage HealDirect.
func NewHealDirect(c ability.HealDirect) HealDirect {
	return HealDirect(c)
}

// NewDamageDirect converts a domain DamageDirect into a storage DamageDirect.
func NewDamageDirect(c ability.DamageDirect) DamageDirect {
	return DamageDirect(c)
}

// NewHealOverTime converts a domain HealOverTime into a storage HealOverTime.
func NewHealOverTime(c ability.HealOverTime) HealOverTime {
	return HealOverTime(c)
}

// NewDamageOverTime converts a domain DamageOverTime into a storage DamageOverTime.
func NewDamageOverTime(c ability.DamageOverTime) DamageOverTime {
	return DamageOverTime(c)
}

// Domain converts a storage HealDirect into a domain HealDirect.
func (c HealDirect) Domain() ability.HealDirect {
	return ability.HealDirect(c)
}

// Domain converts a storage DamageDirect into a domain DamageDirect.
func (c DamageDirect) Domain() ability.DamageDirect {
	return ability.DamageDirect(c)
}

// Domain converts a storage HealOverTime into a domain HealOverTime.
func (c HealOverTime) Domain() ability.HealOverTime {
	return ability.HealOverTime(c)
}

// Domain converts a storage DamageOverTime into a domain DamageOverTime.
func (c DamageOverTime) Domain() ability.DamageOverTime {
	return ability.DamageOverTime(c)
}

// NewAbility convert a ability.A into a storage Ability.
func NewAbility(a ability.A) *Ability {
	components := make([]interface{}, len(a.Components))
	for i, c := range a.Components {
		switch c.(type) {
		case ability.HealDirect:
			components[i] = NewHealDirect(c.(ability.HealDirect))
		case ability.DamageDirect:
			components[i] = NewDamageDirect(c.(ability.DamageDirect))
		case ability.HealOverTime:
			components[i] = NewHealOverTime(c.(ability.HealOverTime))
		case ability.DamageOverTime:
			components[i] = NewDamageOverTime(c.(ability.DamageOverTime))
		}
	}
	return &Ability{
		ID:            [16]byte(a.ID),
		Type:          [16]byte(a.Type),
		Name:          a.Name,
		MPConsumption: a.MPConsumption,
		CD:            a.CD,
		CurrentCD:     a.CurrentCD,
		Components:    components,
	}
}

// Domain converts a storage Ability into a game Ability.
func (a Ability) Domain() ability.A {
	components := make([]ability.AComponent, len(a.Components))
	for i, c := range a.Components {
		switch c.(type) {
		case HealDirect:
			components[i] = c.(HealDirect).Domain()
		case DamageDirect:
			components[i] = c.(DamageDirect).Domain()
		case HealOverTime:
			components[i] = c.(HealOverTime).Domain()
		case DamageOverTime:
			components[i] = c.(DamageOverTime).Domain()
		}
	}
	return ability.A{
		ID:            game.ID(a.ID),
		Type:          ability.Type(a.Type),
		Name:          a.Name,
		MPConsumption: a.MPConsumption,
		CD:            a.CD,
		CurrentCD:     a.CurrentCD,
		Components:    components,
	}
}
