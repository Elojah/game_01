package game

// AbilityComponent is a part of a skill.
type AbilityComponent interface {
	Apply(target Entity) AbilityFeedback
}
