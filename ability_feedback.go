package game

// AbilityFeedback represents the effects a ability had on a target.
type AbilityFeedback struct {
	ID         ID
	AbilityID  ID
	Components []AbilityFeedbackComponent
}

// AbilityFeedbackComponent is the feedback of AbilityComponent.
type AbilityFeedbackComponent interface {
	Affect(*Entity)
}

// HealDirectFeedback is the feedback of a HealDirect casted by entity.
type HealDirectFeedback struct {
	Amount int64
}

// Affect applies fb on entity.
func (fb HealDirectFeedback) Affect(entity *Entity) {

}

// DamageDirectFeedback is the feedback of a DamageDirect casted by entity.
type DamageDirectFeedback struct {
	Amount int64
}

// Affect applies fb on entity.
func (fb DamageDirectFeedback) Affect(entity *Entity) {

}

// HealOverTimeFeedback is the feedback of a HealOverTime casted by entity.
type HealOverTimeFeedback struct {
}

// Affect applies fb on entity.
func (fb HealOverTimeFeedback) Affect(entity *Entity) {

}

// DamageOverTimeFeedback is the feedback of a DamageOverTime casted by entity.
type DamageOverTimeFeedback struct {
}

// Affect applies fb on entity.
func (fb DamageOverTimeFeedback) Affect(entity *Entity) {

}
