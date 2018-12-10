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

type toolTargetsLog struct {
	common
	Route   string
	Targets int
	Time    int64
}

func (expected toolTargetsLog) Equal(actual toolTargetsLog) error {
	if actual.common != expected.common {
		return fmt.Errorf("unexpected log %s", fmt.Sprint(actual.common))
	}
	if actual.Targets != expected.Targets {
		return fmt.Errorf("invalid targets %d", actual.Targets)
	}
	if actual.Route != expected.Route {
		return fmt.Errorf("invalid route %s", actual.Route)
	}
	return nil
}

type toolMoveSuccessLog struct {
	common
	Route  string
	Entity string
	Time   int64
}

func (expected toolMoveSuccessLog) Equal(actual toolMoveSuccessLog) error {
	if actual.common != expected.common {
		return fmt.Errorf("unexpected log %s", fmt.Sprint(actual.common))
	}
	if _, err := gulid.Parse(actual.Entity); err != nil {
		return fmt.Errorf("invalid entity %s", actual.Entity)
	}
	if actual.Route != expected.Route {
		return fmt.Errorf("invalid route %s", actual.Route)
	}
	return nil
}

func expectToolEntityMove(a *LogAnalyzer, ent entity.E) (entity.E, error) {

	// #Force move via tool entity at current sector frontier.
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // nolint: gosec

	pos := geometry.Position{
		Coord:    geometry.Vec3{X: 1024, Y: 1024, Z: 1024},
		SectorID: gulid.MustParse("01CF001HTBA3CDR1ERJ6RF183A"),
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

	expectedTtsLog := toolTargetsLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_tool",
			Message: "found",
		},
		Targets: 1,
		Route:   "/entity/move",
	}
	expectedTmscLog := toolMoveSuccessLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_tool",
			Message: "tool move success",
		},
		Route: "/entity/move",
	}
	n := 2
	if err := a.Expect(func(s string) (bool, error) {
		var c common
		if err := json.Unmarshal([]byte(s), &c); err != nil {
			return false, err
		}
		switch c.Exe {
		case "./bin/game_tool":
			n--
			switch n {
			case 1:
				var ttsActual toolTargetsLog
				if err := json.Unmarshal([]byte(s), &ttsActual); err != nil {
					return true, err
				}
				return false, expectedTtsLog.Equal(ttsActual)
			case 0:
				var tmscActual toolMoveSuccessLog
				if err := json.Unmarshal([]byte(s), &tmscActual); err != nil {
					return true, err
				}
				return true, expectedTmscLog.Equal(tmscActual)
			}
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
