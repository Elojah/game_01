package main

import (
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

/*
#Test plan
- SUCCESS Move same sector
- FAIL Move same sector too far
- SUCCESS Move neighbour sector
- FAIL Move not neighbour sector
- FAIL Move neighbour sector too far
*/

func expectAPI(a *LogAnalyzer, tok account.Token, ent entity.E) error {
	var c client.C
	c.Dial(client.Config{
		Middlewares: []string{"lz4"},
		PacketSize:  1024,
	})

	// #SUCCESS Move same sector
	moveSameSector := event.DTO{
		ID:    ulid.NewID(),
		Token: tok.ID,
		TS:    time.Now(),
		Query: event.Query{
			Move: &event.Move{
				Source:  ent.ID,
				Targets: []ulid.ID{ent.ID},
				Position: geometry.Position{
					SectorID: ulid.MustParse("01CF001HTBA3CDR1ERJ6RF183A"),
					// TODO ?
					Coord: geometry.Vec3{},
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

	return nil
}
