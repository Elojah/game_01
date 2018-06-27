package storage

import (
	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/ulid"
)

// NewHealDirectFeedback converts a game HealDirectFeedback into a storage HealDirectFeedback.
func NewHealDirectFeedback(fb ability.HealDirectFeedback) HealDirectFeedback {
	return HealDirectFeedback(fb)
}

// NewDamageDirectFeedback converts a game DamageDirectFeedback into a storage DamageDirectFeedback.
func NewDamageDirectFeedback(fb ability.DamageDirectFeedback) DamageDirectFeedback {
	return DamageDirectFeedback(fb)
}

// NewHealOverTimeFeedback converts a game HealOverTimeFeedback into a storage HealOverTimeFeedback.
func NewHealOverTimeFeedback(fb ability.HealOverTimeFeedback) HealOverTimeFeedback {
	return HealOverTimeFeedback(fb)
}

// NewDamageOverTimeFeedback converts a game DamageOverTimeFeedback into a storage DamageOverTimeFeedback.
func NewDamageOverTimeFeedback(fb ability.DamageOverTimeFeedback) DamageOverTimeFeedback {
	return DamageOverTimeFeedback(fb)
}

// Domain converts a HealDirectFeedback storage into a HealDirectFeedback domain.
func (fb HealDirectFeedback) Domain() ability.HealDirectFeedback {
	return ability.HealDirectFeedback(fb)
}

// Domain converts a DamageDirectFeedback storage into a DamageDirectFeedback domain.
func (fb DamageDirectFeedback) Domain() ability.DamageDirectFeedback {
	return ability.DamageDirectFeedback(fb)
}

// Domain converts a HealOverTimeFeedback storage into a HealOverTimeFeedback domain.
func (fb HealOverTimeFeedback) Domain() ability.HealOverTimeFeedback {
	return ability.HealOverTimeFeedback(fb)
}

// Domain converts a DamageOverTimeFeedback storage into a DamageOverTimeFeedback domain.
func (fb DamageOverTimeFeedback) Domain() ability.DamageOverTimeFeedback {
	return ability.DamageOverTimeFeedback(fb)
}

// NewAbilityFeedback converts a game AbilityFeedback into a storage AbilityFeedback.
func NewAbilityFeedback(fb ability.Feedback) *AbilityFeedback {
	components := make([]interface{}, len(fb.Components))
	for i, c := range fb.Components {
		switch c.(type) {
		case ability.HealDirectFeedback:
			components[i] = NewHealDirectFeedback(c.(ability.HealDirectFeedback))
		case ability.DamageDirectFeedback:
			components[i] = NewDamageDirectFeedback(c.(ability.DamageDirectFeedback))
		case ability.HealOverTimeFeedback:
			components[i] = NewHealOverTimeFeedback(c.(ability.HealOverTimeFeedback))
		case ability.DamageOverTimeFeedback:
			components[i] = NewDamageOverTimeFeedback(c.(ability.DamageOverTimeFeedback))
		}
	}
	return &AbilityFeedback{
		ID:         [16]byte(fb.ID),
		AbilityID:  [16]byte(fb.AbilityID),
		Components: components,
	}
}

// Domain converts a storage AbilityFeedback into a game AbilityFeedback.
func (a AbilityFeedback) Domain() ability.Feedback {
	components := make([]ability.FeedbackComponent, len(a.Components))
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
	return ability.Feedback{
		ID:         ulid.ID(a.ID),
		AbilityID:  ulid.ID(a.AbilityID),
		Components: components,
	}
}
