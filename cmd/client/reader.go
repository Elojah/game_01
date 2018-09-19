package main

import (
	"bufio"
	"encoding/json"
	"net"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/mux/client"
)

type reader struct {
	*client.C

	logger zerolog.Logger
	*bufio.Scanner

	addr net.Addr

	ticker    *time.Ticker
	tolerance int64

	events map[ulid.ID]event.DTO
	ack    <-chan ulid.ID
	event  chan event.DTO
}

func newReader(c *client.C, ack <-chan ulid.ID) *reader {
	return &reader{
		C:       c,
		logger:  log.With().Str("app", "reader").Logger(),
		Scanner: bufio.NewScanner(os.Stdin),
		event:   make(chan event.DTO),
		events:  make(map[ulid.ID]event.DTO, 0),
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
	r.tolerance = cfg.Tolerance
	var err error
	if r.addr, err = net.ResolveUDPAddr("udp", cfg.Address); err != nil {
		return err
	}

	d := time.Duration(r.tolerance * 1000)
	r.ticker = time.NewTicker(d)
	go r.Run()
	go r.HandleACK()
	return nil
}

// Run starts to read JSON data from stdin and sends it to API.
func (r reader) Run() {
	for r.Scan() {
		var input event.DTO
		if err := json.Unmarshal(r.Scanner.Bytes(), &input); err != nil {
			r.logger.Error().Err(err).Msg("failed to decode input")
			continue
		}
		raw, err := input.Marshal()
		if err != nil {
			r.logger.Error().Err(err).Msg("failed to marshal action")
			continue
		}
		r.event <- input
		go r.Send(raw, r.addr)
	}
}

// HandleACK handles events sending and received acks.
func (r reader) HandleACK() {
	for {
		select {
		case <-r.ticker.C:
			now := time.Now().Unix()
			for _, e := range r.events {
				t := e.ID.Time()
				if t > now || now-t < r.tolerance {
					continue
				}
				go func(e event.DTO) {
					raw, err := e.Marshal()
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
