package ability

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// Template alias an ability.A.
// It represents semi static data. When creating abilities, those templates are used.
type Template = A

// TemplateStore is the communication interface for ability templates.
type TemplateStore interface {
	InsertTemplate(Template) error
	FetchTemplate(ulid.ID) (Template, error)
	ListTemplate() ([]Template, error)
}
