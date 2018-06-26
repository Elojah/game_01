package ability

import (
	"github.com/elojah/game_01/pkg/entity"
)

// FeedbackComponent is the feedback of AbilityComponent.
type FeedbackComponent interface {
	Affect(*entity.E)
}

// HealDirectFeedback is the feedback of a HealDirect casted by entity.
type HealDirectFeedback struct {
	Amount int64
}

// Affect applies fb on entity.
func (fb HealDirectFeedback) Affect(e *entity.E) {

}

// DamageDirectFeedback is the feedback of a DamageDirect casted by entity.
type DamageDirectFeedback struct {
	Amount int64
}

// Affect applies fb on entity.
func (fb DamageDirectFeedback) Affect(e *entity.E) {

}

// HealOverTimeFeedback is the feedback of a HealOverTime casted by entity.
type HealOverTimeFeedback struct {
}

// Affect applies fb on entity.
func (fb HealOverTimeFeedback) Affect(e *entity.E) {

}

// DamageOverTimeFeedback is the feedback of a DamageOverTime casted by entity.
type DamageOverTimeFeedback struct {
}

// Affect applies fb on entity.
func (fb DamageOverTimeFeedback) Affect(e *entity.E) {

}
