package tool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/geometry"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// EntityMove force moves an entity without distance checking.
func (s *Service) EntityMove(id gulid.ID, pos geometry.Position) error {

	raw, err := json.Marshal(event.Move{
		Targets:  []gulid.ID{id},
		Position: pos,
	})
	if err != nil {
		return err
	}
	resp, err := http.Post("https://localhost:8081/entity/move", "application/json", bytes.NewReader(raw))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code %d", resp.StatusCode)
	}
	return nil
}
