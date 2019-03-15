package client

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// Loot loots an item from entity.
func (s *Service) Loot(tokID gulid.ID, loot event.Loot) error {

	lootDTO := event.DTO{
		ID:    gulid.NewID(),
		Token: tokID,
		Query: event.Query{
			Loot: &loot,
		},
	}
	raw, err := json.Marshal(lootDTO)
	raw = append(raw, '\n')
	if err != nil {
		return errors.Wrap(fmt.Errorf("failed to marshal payload"), "loot")
	}

	if _, err := io.WriteString(s.LA.Processes["client"].In, string(raw)); err != nil {
		return errors.Wrap(err, "move same sector")
	}

	return nil
}
