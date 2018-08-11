package entity

import (
	"time"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/game_01/pkg/usecase/listener"

	"github.com/pkg/errors"
)

// Move moves entity to position p.
func (e *E) Move(p geometry.Vec3) {
	e.Position.Coord = p
}

// Store is an interface for E object.
type Store interface {
	SetEntity(E, int64) error
	GetEntity(Subset) (E, error)
	DelEntity(Subset) error
}

// Subset is a subset to retrieve one entity.
type Subset struct {
	ID     ulid.ID
	MaxTS  int64
	Cursor uint64
	Count  int64
}

// Service represents entity usecases.
type Service struct {
	EntityStore         Store
	PermissionStore     PermissionStore
	SectorEntitiesStore sector.EntitiesStore

	ListenerService listener.Service
}

// Disconnect disconnects an entity.
func (s Service) Disconnect(id ulid.ID, tok account.Token) error {

	e, err := s.EntityStore.GetEntity(entity.Subset{ID: id, MaxTS: time.Now().UnixNano()})
	if err != nil {
		return errors.Wrapf(err, "get entity %s", id.String())
	}

	// #Close entity listener
	if err := s.ListenerService.Remove(id); err != nil {
		return errors.Wrapf(err, "delete listener %s", id.String())
	}

	// #Delete token permission on entity.
	if err := s.PermissionStore.DelPermission(entity.PermissionSubset{
		Source: tok.ID.String(),
		Target: id.String(),
	}); err != nil {
		return errors.Wrapf(err, "delete permission %s %s", tok.ID.String(), id.String())
	}

	// #Delete pc entity position
	if err := s.SectorEntitiesStore.RemoveEntityToSector(id, e.Position.SectorID); err != nil {
		return errors.Wrapf(err, "remove entity %s to sector %s", id.String(), e.Position.SectorID.String())
	}

	// #Delete pc entity
	if err := s.EntityStore.DelEntity(entity.Subset{ID: id}); err != nil {
		return errors.Wrapf(err, "delete entity %s", id.String())
	}

	return nil
}
