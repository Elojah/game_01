package client

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// Consume consumes an item in entity inventory.
func (s *Service) Consume(tokID gulid.ID, consume event.Consume) error {

	consumeDTO := event.DTO{
		ID:    gulid.NewID(),
		Token: tokID,
		Query: event.Query{
			Consume: &consume,
		},
	}
	raw, err := json.Marshal(consumeDTO)
	raw = append(raw, '\n')
	if err != nil {
		return errors.Wrap(fmt.Errorf("failed to marshal payload"), "consume")
	}

	if _, err := io.WriteString(s.LA.Processes["client"].In, string(raw)); err != nil {
		return errors.Wrap(err, "move same sector")
	}

	return nil
}
