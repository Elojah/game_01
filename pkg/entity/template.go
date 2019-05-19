package entity

import "github.com/elojah/game_01/pkg/ulid"

// Template alias an entity.
// It represents semi static data. When creating PC/Entities, those templates are used.
type Template = E

// TemplateStore contains basic operations fo entity Template object.
type TemplateStore interface {
	InsertTemplate(Template) error
	FetchTemplate(ulid.ID) (Template, error)
	ListTemplate() ([]Template, error)
}
