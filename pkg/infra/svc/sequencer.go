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
		return infra.Sequencer{}, errors.Wrap(err, "get random core")
	}
	seq := infra.Sequencer{ID: id, Action: infra.Open, Pool: c.ID}
	if err := s.InfraQSequencer.PublishSequencer(seq, c.ID); err != nil {
		return infra.Sequencer{}, errors.Wrapf(err, "open sequencer %s on core %s", seq.ID.String(), c.ID.String())
	}

	// #Set sequencer with saved core id
	if err := s.InfraSequencer.SetSequencer(seq); err != nil {
		return infra.Sequencer{}, errors.Wrapf(err, "set sequencer %s", seq.ID)
	}
	return seq, nil
}

// Remove deletes a sequencer id on any pool.
func (s *SequencerService) Remove(id ulid.ID) error {

	// #Retrieve and close sequencer
	seq, err := s.InfraSequencer.GetSequencer(id)
	if err != nil {
		return errors.Wrapf(err, "get sequencer %s", id.String())
	}
	seq.Action = infra.Close
	if err := s.InfraQSequencer.PublishSequencer(seq, seq.Pool); err != nil {
		return errors.Wrapf(err, "close sequencer %s on pool %s", seq.ID.String(), seq.Pool.String())
	}

	// #Delete sequencer
	if err := s.InfraSequencer.DelSequencer(seq.ID); err != nil {
		return errors.Wrapf(err, "delete sequencer %s", seq.ID.String())
	}
	return nil
}
