package tool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	gulid "github.com/elojah/game_01/pkg/ulid"
)

// Loot set entities as lootable.
func (s *Service) Loot(ids []gulid.ID) error {

	raw, err := json.Marshal(ids)
	if err != nil {
		return err
	}
	resp, err := http.Post(s.url+"/loot", "application/json", bytes.NewReader(raw))
	if err != nil {
		return err
	}
	_ = resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code %d", resp.StatusCode)
	}
	return nil
}
