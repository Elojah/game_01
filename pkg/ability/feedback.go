package ability

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// FeedbackStore contains basic operations for ability feedback object.
type FeedbackStore interface {
	InsertFeedback(Feedback) error
	FetchFeedback(ulid.ID) (Feedback, error)
}
