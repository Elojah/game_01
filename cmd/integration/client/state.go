package client

import (
	"encoding/json"
	"fmt"

	"github.com/elojah/game_01/pkg/entity"
	gerrors "github.com/elojah/game_01/pkg/errors"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

func (s *Service) GetState(id gulid.ID) (entity.E, error) {

	var actual entity.E
	// entLog, err :=
	if err := json.Unmarshal([]byte(s), &actual); err != nil {
		return false, fmt.Errorf("invalid entity %s", s)
	}
	if actual.ID.IsZero() {
		// wrong log type (packet processed)
		return false, nil
	}
	if actual.ID.Compare(ent.ID) == 0 &&
		actual.Position.SectorID.Compare(ent.Position.SectorID) == 0 &&
		actual.Position.Coord == newCoord {
		return true, nil
	}
	return false, nil

	return gerrors.ErrNotImplementedYet{}
}
