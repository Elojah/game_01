package main

import (
	"bufio"
	"encoding/json"
	"net"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/dto"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/mux/client"
)

type reader struct {
	*client.C

	logger zerolog.Logger
	*bufio.Scanner

	token ulid.ID
	addr  net.Addr

	ticker    *time.Ticker
	tolerance time.Duration

	events map[ulid.ID]dto.Event
	ack    <-chan ulid.ID
	event  chan dto.Event
}

func newReader(c *client.C, ack <-chan ulid.ID) *reader {
	return &reader{
		C:       c,
		logger:  log.With().Str("app", "reader").Logger(),
		Scanner: bufio.NewScanner(os.Stdin),
		event:   make(chan dto.Event),
		ack:     ack,
	}
}

func (r *reader) Close() error {
	if err := r.C.Close(); err != nil {
		return err
	}
	close(r.event)
	r.ticker.Stop()
	return nil
}

// Dial initialize a reader.
func (r *reader) Dial(cfg Config) error {
	r.token = cfg.Token
	r.tolerance = cfg.Tolerance
	var err error
	if r.addr, err = net.ResolveUDPAddr("udp", cfg.Address); err != nil {
		return err
	}

	r.ticker = time.NewTicker(r.tolerance)
	go r.Start()
	go r.HandleACK()
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
		e := dto.Event{
			ID:     ulid.NewID(),
			Token:  r.token,
			TS:     time.Now().UnixNano(),
			Action: input.Action,
		}
		raw, err := e.Marshal(nil)
		if err != nil {
			r.logger.Error().Err(err).Msg("failed to marshal action")
			continue
		}
		r.event <- e
		go r.Send(raw, r.addr)
	}
}

// HandleACK handles events sending and received acks.
func (r reader) HandleACK() {
	for {
		select {
		case <-r.ticker.C:
			now := time.Now()
			for _, e := range r.events {
				if now.Sub(time.Unix(0, e.TS)) < r.tolerance {
					continue
				}
				go func(e dto.Event) {
					raw, err := e.Marshal(nil)
					if err != nil {
						r.logger.Error().Err(err).Msg("failed to marshal action")
						return
					}
					r.Send(raw, r.addr)
				}(e)
			}
		case e := <-r.event:
			r.events[e.ID] = e
		case id := <-r.ack:
			delete(r.events, id)
		}
	}
}
