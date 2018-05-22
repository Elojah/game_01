package game

// AbilityTemplate alias an entity.
// It represents semi static data. When creating abilitys, those templates are used.
type AbilityTemplate Ability

// AbilityTemplateMapper is an interface for AbilityTemplate object.
type AbilityTemplateMapper interface {
	SetAbilityTemplate(AbilityTemplate) error
	GetAbilityTemplate(AbilityTemplateSubset) (AbilityTemplate, error)
	ListAbilityTemplate() ([]AbilityTemplate, error)
}

// AbilityTemplateSubset is a subset to retrieve one template.
type AbilityTemplateSubset struct {
	Type AbilityType
}
