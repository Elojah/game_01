package infra

import (
	"github.com/elojah/game_01/pkg/ulid"
)

// QSequencerStore handles send/receive methods for sequencers.
type QSequencerStore interface {
	PublishSequencer(Sequencer, ulid.ID) error
	SubscribeSequencer(ulid.ID) *Subscription
}

// SequencerStore handle sequencer data interactions.
type SequencerStore interface {
	SetSequencer(Sequencer) error
	GetSequencer(SequencerSubset) (Sequencer, error)
	DelSequencer(SequencerSubset) error
}

// SequencerSubset retrieves sequencer per ID.
type SequencerSubset struct {
	ID ulid.ID
}

// SequencerService represents sequencer usecases.
type SequencerService interface {
	New(ulid.ID) (Sequencer, error)
	Remove(ulid.ID) error
}
