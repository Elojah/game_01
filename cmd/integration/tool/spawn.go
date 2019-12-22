package tool

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/entity"
)

// AddSpawn add spawns listed in json filename.
func (s *Service) AddSpawn(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return errors.Wrap(err, "add spawns")
	}
	defer file.Close()
	resp, err := http.Post(s.url+"/spawn", "application/json", file)
	if err != nil {
		return errors.Wrap(err, "add spawns")
	}
	_ = resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "add spawns")
	}
	return nil
}

// GetSpawn get spawns by ids.
func (s *Service) GetSpawn(ids []string) ([]entity.Spawn, error) {
	req, err := http.NewRequest("GET", s.url+"/spawn", nil)
	if err != nil {
		return nil, errors.Wrap(err, "get spawn")
	}

	q := req.URL.Query()
	q.Add("ids", strings.Join(ids, ","))
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "get spawn")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "get spawn")
	}

	// #Read body
	var spawns []entity.Spawn
	if err := json.NewDecoder(resp.Body).Decode(&spawns); err != nil {
		return nil, errors.Wrap(err, "get spawn")
	}

	return spawns, nil
}
