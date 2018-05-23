package storage

import (
	"github.com/elojah/game_01"
)

// NewHealDirectFeedback converts a game HealDirectFeedback into a storage HealDirectFeedback.
func NewHealDirectFeedback(fb game.HealDirectFeedback) HealDirectFeedback {
	return HealDirectFeedback(fb)
}

// NewDamageDirectFeedback converts a game DamageDirectFeedback into a storage DamageDirectFeedback.
func NewDamageDirectFeedback(fb game.DamageDirectFeedback) DamageDirectFeedback {
	return DamageDirectFeedback(fb)
}

// NewHealOverTimeFeedback converts a game HealOverTimeFeedback into a storage HealOverTimeFeedback.
func NewHealOverTimeFeedback(fb game.HealOverTimeFeedback) HealOverTimeFeedback {
	return HealOverTimeFeedback(fb)
}

// NewDamageOverTimeFeedback converts a game DamageOverTimeFeedback into a storage DamageOverTimeFeedback.
func NewDamageOverTimeFeedback(fb game.DamageOverTimeFeedback) DamageOverTimeFeedback {
	return DamageOverTimeFeedback(fb)
}

// Domain converts a HealDirectFeedback storage into a HealDirectFeedback domain.
func (fb HealDirectFeedback) Domain() game.HealDirectFeedback {
	return game.HealDirectFeedback(fb)
}

// Domain converts a DamageDirectFeedback storage into a DamageDirectFeedback domain.
func (fb DamageDirectFeedback) Domain() game.DamageDirectFeedback {
	return game.DamageDirectFeedback(fb)
}

// Domain converts a HealOverTimeFeedback storage into a HealOverTimeFeedback domain.
func (fb HealOverTimeFeedback) Domain() game.HealOverTimeFeedback {
	return game.HealOverTimeFeedback(fb)
}

// Domain converts a DamageOverTimeFeedback storage into a DamageOverTimeFeedback domain.
func (fb DamageOverTimeFeedback) Domain() game.DamageOverTimeFeedback {
	return game.DamageOverTimeFeedback(fb)
}

// NewAbilityFeedback converts a game AbilityFeedback into a storage AbilityFeedback.
func NewAbilityFeedback(fb game.AbilityFeedback) *AbilityFeedback {
	components := make([]interface{}, len(fb.Components))
	for i, c := range fb.Components {
		switch c.(type) {
		case game.HealDirectFeedback:
			components[i] = NewHealDirectFeedback(c.(game.HealDirectFeedback))
		case game.DamageDirectFeedback:
			components[i] = NewDamageDirectFeedback(c.(game.DamageDirectFeedback))
		case game.HealOverTimeFeedback:
			components[i] = NewHealOverTimeFeedback(c.(game.HealOverTimeFeedback))
		case game.DamageOverTimeFeedback:
			components[i] = NewDamageOverTimeFeedback(c.(game.DamageOverTimeFeedback))
		}
	}
	return &AbilityFeedback{
		ID:         [16]byte(fb.ID),
		AbilityID:  [16]byte(fb.AbilityID),
		Components: components,
	}
}

// Domain converts a storage AbilityFeedback into a game AbilityFeedback.
func (a AbilityFeedback) Domain() game.AbilityFeedback {
	components := make([]game.AbilityFeedbackComponent, len(a.Components))
	for i, c := range a.Components {
		switch c.(type) {
		case HealDirectFeedback:
			components[i] = c.(HealDirectFeedback).Domain()
		case DamageDirectFeedback:
			components[i] = c.(DamageDirectFeedback).Domain()
		case HealOverTimeFeedback:
			components[i] = c.(HealOverTimeFeedback).Domain()
		case DamageOverTimeFeedback:
			components[i] = c.(DamageOverTimeFeedback).Domain()
		}
	}
	return game.AbilityFeedback{
		ID:         game.ID(a.ID),
		AbilityID:  game.ID(a.AbilityID),
		Components: components,
	}
}
