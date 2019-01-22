package svc

import (
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"

	"github.com/pkg/errors"
)

// SequencerService represents sequencer usecases.
type SequencerService struct {
	InfraQSequencer infra.QSequencerStore
	InfraSequencer  infra.SequencerStore
	InfraCore       infra.CoreStore
}

// New creates a new sequencer on a random core for id id.
func (s *SequencerService) New(id ulid.ID) (infra.Sequencer, error) {

	// #Open sequencer on a random core
	c, err := s.InfraCore.GetRandomCore()
	if err != nil {
		return infra.Sequencer{}, errors.Wrap(err, "new sequencer")
	}
	seq := infra.Sequencer{ID: id, Action: infra.Open, Pool: c.ID}
	if err := s.InfraQSequencer.PublishSequencer(seq, c.ID); err != nil {
		return infra.Sequencer{}, errors.Wrap(err, "new sequencer")
	}

	// #Set sequencer with saved core id
	if err := s.InfraSequencer.SetSequencer(seq); err != nil {
		return infra.Sequencer{}, errors.Wrap(err, "new sequencer")
	}
	return seq, nil
}

// Remove deletes a sequencer id on any pool.
func (s *SequencerService) Remove(id ulid.ID) error {

	// #Retrieve and close sequencer
	seq, err := s.InfraSequencer.GetSequencer(id)
	if err != nil {
		return errors.Wrapf(err, "remove sequencer")
	}
	seq.Action = infra.Close
	if err := s.InfraQSequencer.PublishSequencer(seq, seq.Pool); err != nil {
		return errors.Wrapf(err, "remove sequencer")
	}

	// #Delete sequencer
	if err := s.InfraSequencer.DelSequencer(seq.ID); err != nil {
		return errors.Wrapf(err, "remove sequencer")
	}
	return nil
}
