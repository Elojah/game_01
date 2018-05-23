package game

// AbilityFeedback represents the effects a ability had on a target.
type AbilityFeedback struct {
	AbilityID  ID
	Components []AbilityFeedbackComponent
}

type AbilityFeedbackComponent interface{}

type HealDirect struct{}
type DamageDirect struct{}
type HealOverTime struct{}
type DamageOverTime struct{}
