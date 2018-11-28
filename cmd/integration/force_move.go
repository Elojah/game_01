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

	// if err := a.Expect(func(s string) (bool, error) {
	// 	var c common
	// 	if err := json.Unmarshal([]byte(s), &c); err != nil {
	// 		return false, err
	// 	}
	// 	switch c.Exe {
	// 	case "./bin/game_tool":
	// 		// ignore
	// 	case "./bin/game_sync":
	// 		// ignore
	// 	default:
	// 		return false, fmt.Errorf("unexpected exe %s", c.Exe)
	// 	}
	// 	return false, nil
	// }); err != nil {
	// 	return ent, err
	// }

	return ent, nil
}
