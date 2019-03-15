package tool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elojah/game_01/pkg/entity"
)

// Inventory creates or updates inventories.
func (s *Service) Inventory(ins []entity.Inventory) error {

	raw, err := json.Marshal(ins)
	if err != nil {
		return err
	}
	resp, err := http.Post("https://localhost:8081/inventory", "application/json", bytes.NewReader(raw))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code %d", resp.StatusCode)
	}
	return nil
}
