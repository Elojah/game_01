package tool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elojah/game_01/pkg/item"
)

// Item creates or updates items.
func (s *Service) Item(its []item.I) error {

	raw, err := json.Marshal(its)
	if err != nil {
		return err
	}
	resp, err := http.Post("https://localhost:8081/item", "application/json", bytes.NewReader(raw))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code %d", resp.StatusCode)
	}
	return nil
}
