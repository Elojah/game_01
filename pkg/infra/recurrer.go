package infra

import (
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
)

// QRecurrerStore handles send/receive methods for recurrers.
type QRecurrerStore interface {
	PublishRecurrer(Recurrer, ulid.ID) error
	SubscribeRecurrer(ulid.ID) *Subscription
}

// RecurrerStore handles set/get methods for recurrers.
type RecurrerStore interface {
	SetRecurrer(Recurrer) error
	GetRecurrer(RecurrerSubset) (Recurrer, error)
	DelRecurrer(RecurrerSubset) error
}

// RecurrerSubset retrieves recurrer by Token ID.
type RecurrerSubset struct {
	TokenID ulid.ID
}

// RecurrerService represents recurrer usecases.
type RecurrerService struct {
	QRecurrerService QRecurrerService
	RecurrerService  RecurrerService
	SyncService      SyncService
}

// New creates a new recurrer on a random sync for id id.
func (s *RecurrerService) New(entityID ulid.ID, tokenID ulid.ID) (infra.Recurrer, error) {
	sy, err := s.GetRandomSync(infra.SyncSubset{})
	if err != nil {
		return infra.Recurrer{}, errors.Wrap(err, "get random sync")
	}
	r := infra.Recurrer{
		EntityID: entityID,
		TokenID:  tokenID,
		Action:   infra.Open,
		Pool:     sy.ID,
	}
	if err := s.PublishRecurrer(r, sy.ID); err != nil {
		return infra.Recurrer{}, errors.Wrapf(err, "open recurrer %s on pool %s", entityID.String(), sy.ID.String())
	}
	if err := s.SetRecurrer(r); err != nil {
		return infra.Recurrer{}, errors.Wrapf(err, "set recurrer", r.EntityID.String())
	}
	return recurrer, nil
}

// Remove deletes a recurrer id on any pool.
func (s *RecurrerService) Remove(id ulid.ID) error {
	r, err := s.GetRecurrer(infra.RecurrerSubset{TokenID: id})
	if err != nil {
		return errors.Wrapf(err, "get recurrer %s", id.String())
	}
	r.Action = infra.Close
	if err := s.PublishRecurrer(r, r.Pool); err != nil {
		return errors.Wrapf(err, "close recurrer %s on pool %s", r.EntityID.String(), r.Pool.String())
	}
	if err := s.DelRecurrer(infra.RecurrerSubset{TokenID: r.TokenID}); err != nil {
		return errors.Wrapf(err, "delete recurrer for token %s", r.TokenID.String())
	}
	return nil
}
