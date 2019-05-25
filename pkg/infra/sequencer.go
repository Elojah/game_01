package infra

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// QSequencerStore contains basic queue operations for infra sequencer object.
type QSequencerStore interface {
	PublishSequencer(Sequencer, ulid.ID) error
	SubscribeSequencer(ulid.ID) *Subscription
}

// SequencerStore contains basic operations for infra sequencer object.
type SequencerStore interface {
	UpsertSequencer(Sequencer) error
	FetchSequencer(ulid.ID) (Sequencer, error)
	RemoveSequencer(ulid.ID) error
}

// SequencerApp contains sequencer stores and applications.
type SequencerApp interface {
	QSequencerStore
	SequencerStore
	CoreStore

	Create(ulid.ID) (Sequencer, error)
	Erase(ulid.ID) error
}
