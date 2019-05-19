package ability

import (
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// Store contains basic operations for ability A object.
type Store interface {
	Insert(A, gulid.ID) error
	Fetch(gulid.ID, gulid.ID) (A, error)
	List(gulid.ID) ([]A, error)
	Remove(gulid.ID, gulid.ID) error
}

// App contains ability stores and applications.
type App interface {
	FeedbackStore
	StarterStore
	Store
	TemplateStore

	SetStarters(gulid.ID, gulid.ID) error
	Copy(gulid.ID, gulid.ID) error
}
