package ability

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// Template alias an ability.A.
// It represents semi static data. When creating abilities, those templates are used.
type Template = A

// TemplateStore is an interface for Template object.
type TemplateStore interface {
	SetTemplate(Template) error
	GetTemplate(TemplateSubset) (Template, error)
	ListTemplate() ([]Template, error)
}

// TemplateSubset is a subset to retrieve one template.
type TemplateSubset struct {
	Type ulid.ID
}
