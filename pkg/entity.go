package pkg

import (
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
)

// Entity wraps use cases for entity.
type Entity struct {
	account.TokenMapper

	EntityMapper entity.Mapper

	event.QRecurrerMapper
	event.QListenerMapper

	sector.EntitiesMapper
}

// DelEntity deletes an entity and remove it from its sector.
func (e Entity) DelEntity(id ulid.ID) error {
	return nil
}
