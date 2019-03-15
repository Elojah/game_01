package tool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elojah/game_01/pkg/entity"
)

// Entity creates a new entity (with sector add).
func (s *Service) Entity(es []entity.E) error {

	raw, err := json.Marshal(es)
	if err != nil {
		return err
	}
	resp, err := http.Post("https://localhost:8081/entity", "application/json", bytes.NewReader(raw))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code %d", resp.StatusCode)
	}
	return nil
}
