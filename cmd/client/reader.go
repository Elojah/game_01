package main

import (
	"encoding/json"
	"net"
	"os"
	"time"

	"github.com/rs/zerolog"

	game "github.com/elojah/game_01"
	"github.com/elojah/game_01/dto"
	"github.com/elojah/mux"
)

type reader struct {
	*mux.M
	*json.Decoder

	logger zerolog.Logger

	tickRate uint
	token    game.ID
	addr     net.Addr
}

// Dial initialize a reader.
func (r reader) Dial(cfg Config) error {
	r.token = cfg.Token
	r.Decoder = json.NewDecoder(os.Stdin)
	r.tickRate = cfg.TickRate
	var err error
	if r.addr, err = net.ResolveUDPAddr("udp", cfg.Address); err != nil {
		return err
	}
	return nil
}

// Start starts to read JSON data from stdin and sends it to API.
func (r reader) Start() {
	for {
		var input Input
		if err := r.Decode(&input); err != nil {
			r.logger.Error().Err(err).Msg("failed to decode input")
			continue
		}
		message := dto.Message{
			ID:     game.NewID(),
			Token:  r.token,
			TS:     time.Now().UnixNano(),
			Action: input,
		}
		raw, err := message.Marshal(nil)
		if err != nil {
			r.logger.Error().Err(err).Msg("failed to marshal action")
			continue
		}
		go r.Send(raw, r.addr)
	}
}
