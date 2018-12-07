package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/entity"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// abilityWithEntity represents the payload to create/associate new ability.
type abilityWithEntity struct {
	ability.A
	EntityID gulid.ID `json:"entity_id"`
}

func expectToolSetAbility(a *LogAnalyzer, ab ability.A, ent entity.E) error {

	// #Force move via tool entity at current sector frontier.
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // nolint: gosec

	raw, err := json.Marshal([]abilityWithEntity{
		abilityWithEntity{
			EntityID: ent.ID,
			A:        ab,
		},
	})
	if err != nil {
		return err
	}
	resp, err := http.Post("https://localhost:8081/ability", "application/json", bytes.NewReader(raw))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code %d", resp.StatusCode)
	}

	expectedTmscLog := toolMoveSuccessLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_tool",
			Message: "tool move success",
		},
		Route: "/ability",
	}
	if err := a.Expect(func(s string) (bool, error) {
		var c common
		if err := json.Unmarshal([]byte(s), &c); err != nil {
			return false, err
		}
		switch c.Exe {
		case "./bin/game_tool":
			// ignore
			var tmscActual toolMoveSuccessLog
			if err := json.Unmarshal([]byte(s), &tmscActual); err != nil {
				return true, err
			}
			return true, expectedTmscLog.Equal(tmscActual)
		case "./bin/game_sync":
			// ignore
		default:
			return false, fmt.Errorf("unexpected exe %s", c.Exe)
		}
		return false, nil
	}); err != nil {
		return err
	}

	return nil
}
