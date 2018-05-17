package game

// Template alias an entity.
// It represents semi static data. When creating PC/Entities, those templates are used.
type Template Entity

// TemplateService is an interface for Template object.
type TemplateService interface {
	CreateTemplate(Template) error
	GetTemplate(TemplateSubset) (Template, error)
}

// TemplateSubset is a subset to retrieve one template.
type TemplateSubset struct {
	Type string
}
