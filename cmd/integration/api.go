package main

import (
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
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

	return nil
}
