package main

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
	"github.com/oklog/ulid"
)

func expectCast(a *LogAnalyzer, ac *LogAnalyzer, tok account.Token, ab ability.A, ent entity.E, o entity.E) (entity.E, error) {

	// #SUCCESS Cast same sector

	now := ulid.Now()
	cast := event.DTO{
		ID:    gulid.NewTimeID(now),
		Token: tok.ID,
		Query: event.Query{
			Cast: &event.Cast{
				Source:    ent.ID,
				AbilityID: ab.ID,
				Targets: map[string]ability.Targets{
					"01CYBX32YGJ4A4T4SAMMQKQS1H": ability.Targets{
						Entities: []gulid.ID{o.ID},
					},
				},
			},
		},
	}
	raw, err := json.Marshal(cast)
	raw = append(raw, '\n')
	if err != nil {
		return ent, fmt.Errorf("failed to marshal payload")
	}

	if _, err := io.WriteString(ac.Processes["client"].In, string(raw)); err != nil {
		return ent, err
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
		Action: "cast",
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
		Event: cast.ID,
	}
	expectedAPYLog := applyLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_core",
			Message: "apply",
		},
	}
	expectedAPDLog := appliedLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_core",
			Message: "applied",
		},
		Type: "cast_source",
	}

	nAPI := 3
	nCore := 4
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
			case 3:
				var erActual eventReceivedLog
				if err := json.Unmarshal([]byte(s), &erActual); err != nil {
					return nAPI == 0 && nCore == 0, err
				}
				return nAPI == 0 && nCore == 0, expectedERLog.Equal(erActual)
			case 2:
				var feActual fetchEventLog
				if err := json.Unmarshal([]byte(s), &feActual); err != nil {
					return nAPI == 0 && nCore == 0, err
				}
				return nAPI == 0 && nCore == 0, expectedFELog.Equal(feActual)
			case 1:
				var apyActual applyLog
				if err := json.Unmarshal([]byte(s), &apyActual); err != nil {
					return nAPI == 0 && nCore == 0, err
				}
				return nAPI == 0 && nCore == 0, expectedAPYLog.Equal(apyActual)
			case 0:
				var apdActual appliedLog
				if err := json.Unmarshal([]byte(s), &apdActual); err != nil {
					return nAPI == 0 && nCore == 0, err
				}
				return nAPI == 0 && nCore == 0, expectedAPDLog.Equal(apdActual)
			}
			return nAPI == 0 && nCore == 0, nil
		case "./bin/game_sync":
			// ignore
		case "./bin/game_tool":
			// ignore
		default:
			return false, fmt.Errorf("unexpected exe %s", c.Exe)
		}
		return false, nil
	}); err != nil {
		return ent, err
	}

	// Check new position received and echoed by client.
	tolerance := 200 * time.Millisecond
	timer := time.NewTimer(tolerance)

	var actual entity.E
	defer timer.Stop()
	return actual, ac.Expect(func(s string) (bool, error) {
		select {
		case <-timer.C:
			return false, fmt.Errorf("cast not applied in %s", tolerance.String())
		default:
		}
		if err := json.Unmarshal([]byte(s), &actual); err != nil {
			return false, fmt.Errorf("invalid entity %s", s)
		}
		if actual.ID.IsZero() {
			// wrong log type (packet processed)
			return false, nil
		}
		if actual.ID.Compare(ent.ID) == 0 &&
			actual.Position.SectorID.Compare(ent.Position.SectorID) == 0 &&
			actual.Position.Coord == ent.Position.Coord &&
			actual.MP == ent.MP-ab.MPConsumption &&
			actual.Cast.AbilityID.Compare(ab.ID) == 0 {
			return true, nil
		}
		return false, nil
	})
}