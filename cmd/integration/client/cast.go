package client

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// Cast move an entity at position pos. Don't check any distance limit reach.
func (s *Service) Cast(tokID gulid.ID, cast event.Cast) error {

	castDTO := event.DTO{
		ID:    gulid.NewID(),
		Token: tokID,
		Query: event.Query{
			Cast: &cast,
		},
	}
	raw, err := json.Marshal(castDTO)
	raw = append(raw, '\n')
	if err != nil {
		return errors.Wrap(fmt.Errorf("failed to marshal payload"), "cast")
	}

	if _, err := io.WriteString(s.LA.Processes["client"].In, string(raw)); err != nil {
		return errors.Wrap(err, "move same sector")
	}

	return nil
}
