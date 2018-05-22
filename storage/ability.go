package storage

import (
	"github.com/elojah/game_01"
)

// HealDirect
// DamageDirect
// HealOverTime
// DamageOverTime

// NewAbility convert a game.Ability into a storage Ability.
func NewAbility(a game.Ability) *Ability {
	return &Ability{
		ID:            [16]byte(a.ID),
		Type:          [16]byte(a.Type),
		Name:          a.Name,
		MPConsumption: a.MPConsumption,
		CD:            a.CD,
		CurrentCD:     a.CurrentCD,
	}
}

// Domain converts a storage Ability into a game Ability.
func (a Ability) Domain() game.Ability {
	return game.Ability{
		ID:            game.ID(a.ID),
		Type:          game.AbilityType(a.Type),
		Name:          a.Name,
		MPConsumption: a.MPConsumption,
		CD:            a.CD,
		CurrentCD:     a.CurrentCD,
	}
}
