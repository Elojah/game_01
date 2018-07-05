package main

import (
	"bufio"
	"encoding/json"
	"net"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/dto"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/mux/client"
)

type reader struct {
	*client.C

	logger zerolog.Logger
	*bufio.Scanner

	token ulid.ID
	addr  net.Addr
}

func newReader(c *client.C) *reader {
	return &reader{
		C:       c,
		logger:  log.With().Str("app", "reader").Logger(),
		Scanner: bufio.NewScanner(os.Stdin),
	}
}

// Dial initialize a reader.
func (r *reader) Dial(cfg Config) error {
	r.token = cfg.Token
	var err error
	if r.addr, err = net.ResolveUDPAddr("udp", cfg.Address); err != nil {
		return err
	}
	go r.Start()
	return nil
}

// Start starts to read JSON data from stdin and sends it to API.
func (r reader) Start() {
	for r.Scan() {
		var input Input
		if err := json.Unmarshal(r.Scanner.Bytes(), &input); err != nil {
			r.logger.Error().Err(err).Msg("failed to decode input")
			continue
		}
		message := dto.Message{
			ID:     ulid.NewID(),
			Token:  r.token,
			TS:     time.Now().UnixNano(),
			Action: input.Action,
		}
		raw, err := message.Marshal(nil)
		if err != nil {
			r.logger.Error().Err(err).Msg("failed to marshal action")
			continue
		}
		go r.Send(raw, r.addr)
	}
}
