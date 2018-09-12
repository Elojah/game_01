package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/mux/client"
)

type packetProcLog struct {
	common
	Packet string
	Addr   string
	Status string
	Time   int64
}

func (expected packetProcLog) Equal(actual packetProcLog) error {
	if actual.common != expected.common {
		return fmt.Errorf("unexpected log %s", fmt.Sprint(actual.common))
	}
	if _, err := ulid.Parse(actual.Packet); err != nil {
		return fmt.Errorf("invalid packet %s", actual.Packet)
	}
	if actual.Status != expected.Status {
		return fmt.Errorf("invalid status %s", actual.Status)
	}
	if _, err := net.ResolveUDPAddr("udp", actual.Addr); err != nil {
		return fmt.Errorf("invalid addr %s", actual.Addr)
	}
	return nil
}

type packetSentLog struct {
	common
	Bytes   int
	Address string
	Time    int64
}

func (expected packetSentLog) Equal(actual packetSentLog) error {
	if actual.common != expected.common {
		return fmt.Errorf("unexpected log %s", fmt.Sprint(actual.common))
	}
	if actual.Bytes != expected.Bytes {
		return fmt.Errorf("invalid bytes %d", actual.Bytes)
	}
	if _, err := net.ResolveUDPAddr("udp", actual.Address); err != nil {
		return fmt.Errorf("invalid addr %s", actual.Address)
	}
	return nil
}

type eventSendLog struct {
	common
	Packet string
	Action string
	Event  string
	Source string
	Time   int64
}

func (expected eventSendLog) Equal(actual eventSendLog) error {
	if actual.common != expected.common {
		return fmt.Errorf("unexpected log %s", fmt.Sprint(actual.common))
	}
	if _, err := ulid.Parse(actual.Packet); err != nil {
		return fmt.Errorf("invalid packet %s", actual.Packet)
	}
	if actual.Action != expected.Action {
		return fmt.Errorf("invalid action %s", actual.Action)
	}
	if actual.Source != expected.Source {
		return fmt.Errorf("invalid source %s", actual.Source)
	}
	if _, err := ulid.Parse(actual.Event); err != nil {
		return fmt.Errorf("invalid event %s", actual.Event)
	}

	return nil
}

type eventReceivedLog struct {
	common
	Sequencer string
	Event     string
	Time      int64
}

func (expected eventReceivedLog) Equal(actual eventReceivedLog) error {
	if actual.common != expected.common {
		return fmt.Errorf("unexpected log %s", fmt.Sprint(actual.common))
	}
	if _, err := ulid.Parse(actual.Sequencer); err != nil {
		return fmt.Errorf("invalid sequencer %s", actual.Sequencer)
	}
	if _, err := ulid.Parse(actual.Event); err != nil {
		return fmt.Errorf("invalid event %s", actual.Event)
	}

	return nil
}

type fetchEventLog struct {
	common
	Sequencer string
	Current   int64
	Time      int64
}

func (expected fetchEventLog) Equal(actual fetchEventLog) error {
	if actual.common != expected.common {
		return fmt.Errorf("unexpected log %s", fmt.Sprint(actual.common))
	}
	if _, err := ulid.Parse(actual.Sequencer); err != nil {
		return fmt.Errorf("invalid sequencer %s", actual.Sequencer)
	}
	if actual.Current != expected.Current {
		return fmt.Errorf("invalid current %d", actual.Current)
	}

	return nil
}

type applyLog struct {
	common
	Sequencer string
	Event     string
	TS        int64
	Time      int64
}

func (expected applyLog) Equal(actual applyLog) error {
	if actual.common != expected.common {
		return fmt.Errorf("unexpected log %s", fmt.Sprint(actual.common))
	}
	if _, err := ulid.Parse(actual.Sequencer); err != nil {
		return fmt.Errorf("invalid sequencer %s", actual.Sequencer)
	}
	if _, err := ulid.Parse(actual.Event); err != nil {
		return fmt.Errorf("invalid event %s", actual.Event)
	}
	if actual.TS != expected.TS {
		return fmt.Errorf("invalid ts %d", actual.TS)
	}

	return nil
}

type appliedLog struct {
	common
	Sequencer string
	TS        int64
	Type      string
	Time      int64
}

func (expected appliedLog) Equal(actual appliedLog) error {
	if actual.common != expected.common {
		return fmt.Errorf("unexpected log %s", fmt.Sprint(actual.common))
	}

	if _, err := ulid.Parse(actual.Sequencer); err != nil {
		return fmt.Errorf("invalid sequencer %s", actual.Sequencer)
	}
	if actual.TS != expected.TS {
		return fmt.Errorf("invalid ts %d", actual.TS)
	}
	if actual.Type != expected.Type {
		return fmt.Errorf("invalid type %s", actual.Type)
	}

	return nil
}

/*
#Test plan
- SUCCESS Move same sector
- FAIL Move same sector too far
- SUCCESS Move neighbour sector
- FAIL Move not neighbour sector
- FAIL Move neighbour sector too far
*/

func expectMove(a *LogAnalyzer, ac *LogAnalyzer, tok account.Token, ent entity.E) error {
	var c client.C
	c.Dial(client.Config{
		PacketSize: 1024,
	})

	// #SUCCESS Move same sector
	newCoord := geometry.Vec3{
		X: ent.Position.Coord.X + 33,
		Y: ent.Position.Coord.Y + 33,
		Z: ent.Position.Coord.Z + 33,
	}
	if newCoord.X > 1024 {
		newCoord.X = 1024
	}
	if newCoord.Y > 1024 {
		newCoord.Y = 1024
	}
	if newCoord.Z > 1024 {
		newCoord.Z = 1024
	}

	now := time.Now()
	moveSameSector := event.DTO{
		ID:    ulid.NewID(),
		Token: tok.ID,
		TS:    now,
		Query: event.Query{
			Move: &event.Move{
				Source:  ent.ID,
				Targets: []ulid.ID{ent.ID},
				Position: geometry.Position{
					SectorID: ulid.MustParse("01CF001HTBA3CDR1ERJ6RF183A"),
					Coord:    newCoord,
				},
			},
		},
	}
	raw, err := moveSameSector.Marshal()
	if err != nil {
		return fmt.Errorf("failed to marshal payload")
	}
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:3400")
	if err != nil {
		return err
	}
	c.Send(raw, addr)

	nAPI := 3
	nCore := 4
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
		Current: now.UnixNano(),
	}
	expectedAPYLog := applyLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_core",
			Message: "apply",
		},
		TS: now.UnixNano(),
	}
	expectedAPDLog := appliedLog{
		common: common{
			Level:   "info",
			Exe:     "./bin/game_core",
			Message: "applied",
		},
		TS:   now.UnixNano(),
		Type: "move_source",
	}

	if err := a.Expect(func(s string) (bool, error) {
		fmt.Print(s)
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
		default:
			return false, fmt.Errorf("unexpected exe %s", c.Exe)
		}
		// dead code for compile only
		return false, nil
	}); err != nil {
		return err
	}

	tolerance := 200 * time.Millisecond
	timer := time.NewTimer(tolerance)
	defer timer.Stop()
	return ac.Expect(func(s string) (bool, error) {
		select {
		case <-timer.C:
			return false, fmt.Errorf("move not applied in %s", tolerance.String())
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
			actual.Position.Coord == newCoord {
			return true, nil
		}
		fmt.Println("actual:", actual.Position.Coord)
		fmt.Println("newcoord:", newCoord)
		return false, nil
	})
}
