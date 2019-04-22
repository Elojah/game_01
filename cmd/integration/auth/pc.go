package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/elojah/game_01/pkg/entity"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// CreatePC creates a new PC for token.
func (s *Service) CreatePC(tokenID gulid.ID, pcName string, pcType string, spawnID string) error {
	raw, err := json.Marshal(map[string]string{
		"token": tokenID.String(),
		"name":  pcName,
		"type":  pcType,
		"spawn": spawnID,
	})
	if err != nil {
		return errors.Wrap(err, "create pc")
	}
	resp, err := http.Post(s.url+"/pc/create", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return errors.Wrap(err, "create pc")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "create pc")
	}
	return nil
}

// ListPC list all pcs corresponding to token.
func (s *Service) ListPC(tokenID gulid.ID) ([]entity.PC, error) {
	raw, err := json.Marshal(map[string]string{
		"token": tokenID.String(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "list pc")
	}
	resp, err := http.Post(s.url+"/pc/list", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return nil, errors.Wrap(err, "list pc")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "list pc")
	}

	var pcs []entity.PC
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&pcs); err != nil {
		return nil, errors.Wrap(err, "list pc")
	}
	return pcs, nil
}

// DelPC deletes a PC.
func (s *Service) DelPC(tokenID gulid.ID, pcID gulid.ID) error {
	raw, err := json.Marshal(map[string]string{
		"token": tokenID.String(),
		"pc":    pcID.String(),
	})
	if err != nil {
		return errors.Wrap(err, "del pc")
	}
	resp, err := http.Post(s.url+"/pc/del", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return errors.Wrap(err, "del pc")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(fmt.Errorf("invalid status code %d", resp.StatusCode), "create pc")
	}
	return nil
}

// ConnectPC connects to a PC associated to token/account and returns the corresponding entity.
func (s *Service) ConnectPC(tokenID gulid.ID, pcID gulid.ID) (entity.E, error) {
	raw, err := json.Marshal(map[string]string{
		"token":  tokenID.String(),
		"target": pcID.String(),
	})
	if err != nil {
		return entity.E{}, err
	}
	resp, err := http.Post(s.url+"/pc/connect", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return entity.E{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return entity.E{}, fmt.Errorf("invalid status code %d", resp.StatusCode)
	}

	var e entity.E
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
		return entity.E{}, err
	}
	return e, nil
}

// DisconnectPC invalids the token and disconnect pc entity.
func (s *Service) DisconnectPC(tokenID gulid.ID) error {
	raw, err := json.Marshal(map[string]string{
		"token": tokenID.String(),
	})
	if err != nil {
		return err
	}
	resp, err := http.Post(s.url+"/pc/disconnect", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code %d", resp.StatusCode)
	}
	return nil
}
