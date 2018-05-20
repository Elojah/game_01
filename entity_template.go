package game

// EntityTemplate alias an entity.
// It represents semi static data. When creating PC/Entities, those templates are used.
type EntityTemplate Entity

// EntityTemplateService is an interface for EntityTemplate object.
type EntityTemplateService interface {
	SetEntityTemplate(EntityTemplate) error
	GetEntityTemplate(EntityTemplateSubset) (EntityTemplate, error)
}

// EntityTemplateSubset is a subset to retrieve one template.
type EntityTemplateSubset struct {
	Type string
}
