package entity

import "github.com/elojah/game_01/pkg/ulid"

// Template alias an entity.
// It represents semi static data. When creating PC/Entities, those templates are used.
type Template = E

// TemplateStore is an interface for Template object.
type TemplateStore interface {
	SetTemplate(Template) error
	GetTemplate(ulid.ID) (Template, error)
	ListTemplate() ([]Template, error)
}
