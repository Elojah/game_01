package entity

// Template alias an entity.
// It represents semi static data. When creating PC/Entities, those templates are used.
type Template E

// TemplateMapper is an interface for Template object.
type TemplateMapper interface {
	SetTemplate(Template) error
	GetTemplate(TemplateSubset) (Template, error)
	ListTemplate() ([]Template, error)
}

// TemplateSubset is a subset to retrieve one template.
type TemplateSubset struct {
	Type Type
}
