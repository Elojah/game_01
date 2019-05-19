package app

import (
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"

	"github.com/pkg/errors"
)

// SequencerApp represents sequencer usecases.
type SequencerApp struct {
	QSequencerStore
	SequencerStore
	CoreStore
}

// Create creates a new sequencer on a random core for id id.
func (app *SequencerApp) Create(id ulid.ID) (infra.Sequencer, error) {

	// #Open sequencer on a random core
	c, err := app.CoreStore.FetchRandomCore()
	if err != nil {
		return infra.Sequencer{}, errors.Wrap(err, "create sequencer")
	}
	seq := infra.Sequencer{ID: id, Action: infra.Open, Pool: c.ID}
	if err := app.QSequencerStore.PublishSequencer(seq, c.ID); err != nil {
		return infra.Sequencer{}, errors.Wrap(err, "create sequencer")
	}

	// #Set sequencer with saved core id
	if err := app.SequencerStore.InsertSequencer(seq); err != nil {
		return infra.Sequencer{}, errors.Wrap(err, "create sequencer")
	}
	return seq, nil
}

// Erase deletes a sequencer id on any pool.
func (app *SequencerApp) Erase(id ulid.ID) error {

	// #Retrieve and close sequencer
	seq, err := app.SequencerStore.FetchSequencer(id)
	if err != nil {
		return errors.Wrapf(err, "erase sequencer")
	}
	seq.Action = infra.Close
	if err := app.QSequencerStore.PublishSequencer(seq, seq.Pool); err != nil {
		return errors.Wrapf(err, "erase sequencer")
	}

	// #Delete sequencer
	if err := app.SequencerStore.RemoveSequencer(seq.ID); err != nil {
		return errors.Wrapf(err, "erase sequencer")
	}
	return nil
}
