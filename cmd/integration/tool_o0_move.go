package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/geometry"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

func expectToolo0Move(a *LogAnalyzer, ent entity.E) (entity.E, error) {

	// #Force move via tool entity at current sector frontier.
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // nolint: gosec

	pos := geometry.Position{
		Coord:    geometry.Vec3{X: 33 + 10, Y: 1024 - 33 + 10, Z: 1024 - 34 + 10},
		SectorID: gulid.MustParse("01CKQQPVZN5KQC8XC9Q9NK8YXQ"),
	}
	raw, err := json.Marshal(event.MoveSource{
		Targets:  []gulid.ID{ent.ID},
		Position: pos,
	})
	if err != nil {
		return ent, err
	}
	resp, err := http.Post("https://localhost:8081/entity/move", "application/json", bytes.NewReader(raw))
	if err != nil {
		return ent, err
	}
	if resp.StatusCode != http.StatusOK {
		return ent, fmt.Errorf("invalid status code %d", resp.StatusCode)
	}

	ent.Position = pos

	expectedTmscLog := toolMoveSuccessLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_tool",
			Message: "tool move success",
		},
		Route: "/entity/move",
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
		return ent, err
	}

	return ent, nil
}
