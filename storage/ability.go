package storage

import (
	"github.com/elojah/game_01"
)

// NewHealDirect converts a domain HealDirect into a storage HealDirect.
func NewHealDirect(c game.HealDirect) HealDirect {
	return HealDirect(c)
}

// NewDamageDirect converts a domain DamageDirect into a storage DamageDirect.
func NewDamageDirect(c game.DamageDirect) DamageDirect {
	return DamageDirect(c)
}

// NewHealOverTime converts a domain HealOverTime into a storage HealOverTime.
func NewHealOverTime(c game.HealOverTime) HealOverTime {
	return HealOverTime(c)
}

// NewDamageOverTime converts a domain DamageOverTime into a storage DamageOverTime.
func NewDamageOverTime(c game.DamageOverTime) DamageOverTime {
	return DamageOverTime(c)
}

// Domain converts a storage HealDirect into a domain HealDirect.
func (c HealDirect) Domain() game.HealDirect {
	return game.HealDirect(c)
}

// Domain converts a storage DamageDirect into a domain DamageDirect.
func (c DamageDirect) Domain() game.DamageDirect {
	return game.DamageDirect(c)
}

// Domain converts a storage HealOverTime into a domain HealOverTime.
func (c HealOverTime) Domain() game.HealOverTime {
	return game.HealOverTime(c)
}

// Domain converts a storage DamageOverTime into a domain DamageOverTime.
func (c DamageOverTime) Domain() game.DamageOverTime {
	return game.DamageOverTime(c)
}

// NewAbility convert a game.Ability into a storage Ability.
func NewAbility(a game.Ability) *Ability {
	components := make([]interface{}, len(a.Components))
	for i, c := range a.Components {
		switch c.(type) {
		case game.HealDirect:
			components[i] = NewHealDirect(c.(game.HealDirect))
		case game.DamageDirect:
			components[i] = NewDamageDirect(c.(game.DamageDirect))
		case game.HealOverTime:
			components[i] = NewHealOverTime(c.(game.HealOverTime))
		case game.DamageOverTime:
			components[i] = NewDamageOverTime(c.(game.DamageOverTime))
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
func (a Ability) Domain() game.Ability {
	components := make([]game.AbilityComponent, len(a.Components))
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
	return game.Ability{
		ID:            game.ID(a.ID),
		Type:          game.AbilityType(a.Type),
		Name:          a.Name,
		MPConsumption: a.MPConsumption,
		CD:            a.CD,
		CurrentCD:     a.CurrentCD,
		Components:    components,
	}
}
