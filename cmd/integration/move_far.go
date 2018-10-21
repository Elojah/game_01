package main

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/oklog/ulid"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/geometry"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

type invalidActionLog struct {
	common
	Core      string
	Sequencer string
	Event     string
	TS        int64
	Type      string
	// Error     int64
	Time int64
}

func (expected invalidActionLog) Equal(actual invalidActionLog) error {
	if actual.common != expected.common {
		return fmt.Errorf("unexpected log %s", fmt.Sprint(actual.common))
	}
	if _, err := gulid.Parse(actual.Core); err != nil {
		return fmt.Errorf("invalid core %s", actual.Core)
	}
	if _, err := gulid.Parse(actual.Sequencer); err != nil {
		return fmt.Errorf("invalid sequencer %s", actual.Sequencer)
	}
	if _, err := gulid.Parse(actual.Event); err != nil {
		return fmt.Errorf("invalid event %s", actual.Event)
	}
	if actual.Type != expected.Type {
		return fmt.Errorf("invalid type %s", actual.Type)
	}
	return nil
}

func expectMoveSameSectorTooFar(a *LogAnalyzer, ac *LogAnalyzer, tok account.Token, ent entity.E) error {

	// #FAIL Move same sector too far
	newCoord := geometry.Vec3{
		X: ent.Position.Coord.X + 33,
		Y: ent.Position.Coord.Y + 33,
		Z: ent.Position.Coord.Z + 35,
	}
	if newCoord.X > 1024 {
		newCoord.X -= 2 * 33
	}
	if newCoord.Y > 1024 {
		newCoord.Y -= 2 * 33
	}
	if newCoord.Z > 1024 {
		newCoord.Z -= 2 * 35
	}

	now := ulid.Now()
	moveSameSector := event.DTO{
		ID:    gulid.NewTimeID(now),
		Token: tok.ID,
		Query: event.Query{
			Move: &event.Move{
				Source:  ent.ID,
				Targets: []gulid.ID{ent.ID},
				Position: geometry.Position{
					SectorID: gulid.MustParse("01CF001HTBA3CDR1ERJ6RF183A"),
					Coord:    newCoord,
				},
			},
		},
	}
	raw, err := json.Marshal(moveSameSector)
	raw = append(raw, '\n')
	if err != nil {
		return fmt.Errorf("failed to marshal payload")
	}

	if _, err := io.WriteString(ac.Processes["client"].In, string(raw)); err != nil {
		return err
	}

	expectedPPLog := packetProcLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_api",
			Message: "packet processed",
		},
		Status: "processed",
	}
	expectedPSLog := packetSentLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_api",
			Message: "packet sent",
		},
		Bytes: 18,
	}
	expectedESLog := eventSendLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_api",
			Message: "send event",
		},
		Action: "move",
		Source: ent.ID.String(),
	}
	expectedERLog := eventReceivedLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_core",
			Message: "event received",
		},
	}
	expectedFELog := fetchEventLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_core",
			Message: "fetch post events",
		},
		Event: moveSameSector.ID,
	}
	expectedAPYLog := applyLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_core",
			Message: "apply",
		},
	}
	expectedIALog := invalidActionLog{
		common: common{
			Level:   "error",
			Exe:     "./bin/game_core",
			Message: "move target rejected",
		},
		Type: "move_target",
	}
	expectedAPDLog := appliedLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_core",
			Message: "applied",
		},
		Type: "move_source",
	}

	nAPI := 3
	nCore := 8
	if err := a.Expect(func(s string) (bool, error) {
		var c common
		if err := json.Unmarshal([]byte(s), &c); err != nil {
			return false, err
		}
		switch c.Exe {
		case "./bin/game_api":
			nAPI--
			switch nAPI {
			case 2:
				var ppActual packetProcLog
				if err := json.Unmarshal([]byte(s), &ppActual); err != nil {
					return nAPI == 0 && nCore == 0, err
				}
				return nAPI == 0 && nCore == 0, expectedPPLog.Equal(ppActual)
			case 1:
				var psActual packetSentLog
				if err := json.Unmarshal([]byte(s), &psActual); err != nil {
					return nAPI == 0 && nCore == 0, err
				}
				return nAPI == 0 && nCore == 0, expectedPSLog.Equal(psActual)
			case 0:
				var esActual eventSendLog
				if err := json.Unmarshal([]byte(s), &esActual); err != nil {
					return nAPI == 0 && nCore == 0, err
				}
				return nAPI == 0 && nCore == 0, expectedESLog.Equal(esActual)
			}
		case "./bin/game_core":
			nCore--
			switch nCore {
			case 7:
				var erActual eventReceivedLog
				if err := json.Unmarshal([]byte(s), &erActual); err != nil {
					return nAPI == 0 && nCore == 0, err
				}
				return nAPI == 0 && nCore == 0, expectedERLog.Equal(erActual)
			case 6:
				var feActual fetchEventLog
				if err := json.Unmarshal([]byte(s), &feActual); err != nil {
					return nAPI == 0 && nCore == 0, err
				}
				return nAPI == 0 && nCore == 0, expectedFELog.Equal(feActual)
			case 5:
				var apyActual applyLog
				if err := json.Unmarshal([]byte(s), &apyActual); err != nil {
					return nAPI == 0 && nCore == 0, err
				}
				return nAPI == 0 && nCore == 0, expectedAPYLog.Equal(apyActual)
			case 4:
				var apdActual appliedLog
				if err := json.Unmarshal([]byte(s), &apdActual); err != nil {
					return nAPI == 0 && nCore == 0, err
				}
				return nAPI == 0 && nCore == 0, expectedAPDLog.Equal(apdActual)
			case 3:
				var erActual eventReceivedLog
				if err := json.Unmarshal([]byte(s), &erActual); err != nil {
					return nAPI == 0 && nCore == 0, err
				}
				expectedFELog.Event = gulid.MustParse(erActual.Event)
				return nAPI == 0 && nCore == 0, expectedERLog.Equal(erActual)
			case 2:
				var feActual fetchEventLog
				if err := json.Unmarshal([]byte(s), &feActual); err != nil {
					return nAPI == 0 && nCore == 0, err
				}
				// Add one to fetch event because move apply at ts+1
				return nAPI == 0 && nCore == 0, expectedFELog.Equal(feActual)
			case 1:
				var apyActual applyLog
				if err := json.Unmarshal([]byte(s), &apyActual); err != nil {
					return nAPI == 0 && nCore == 0, err
				}
				return nAPI == 0 && nCore == 0, expectedAPYLog.Equal(apyActual)
			case 0:
				var iaActual invalidActionLog
				if err := json.Unmarshal([]byte(s), &iaActual); err != nil {
					return nAPI == 0 && nCore == 0, err
				}
				return nAPI == 0 && nCore == 0, expectedIALog.Equal(iaActual)
			}
			return nAPI == 0 && nCore == 0, nil
		case "./bin/game_sync":
			// ignore
		default:
			return false, fmt.Errorf("unexpected exe %s", c.Exe)
		}
		return false, nil
	}); err != nil {
		return err
	}

	// Check new position received and echoed by client.
	tolerance := 200 * time.Millisecond
	timer := time.NewTimer(tolerance)
	defer timer.Stop()
	return ac.Expect(func(s string) (bool, error) {
		select {
		case <-timer.C:
			return false, fmt.Errorf("sync not received in %s", tolerance.String())
		default:
		}
		var actual entity.E
		if err := json.Unmarshal([]byte(s), &actual); err != nil {
			return false, fmt.Errorf("invalid entity %s", s)
		}
		if actual.ID.IsZero() {
			// wrong log type (packet processed)
			return false, nil
		}
		if actual.ID.Compare(ent.ID) == 0 &&
			actual.Position.SectorID.Compare(ent.Position.SectorID) == 0 &&
			actual.Position.Coord == ent.Position.Coord {
			return true, nil
		}
		return false, nil
	})
}
