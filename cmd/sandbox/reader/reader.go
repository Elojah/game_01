package reader

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/oklog/ulid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/mux/client"
)

type R struct {
	*client.C

	logger zerolog.Logger
	*bufio.Scanner

	addr net.Addr

	ticker    *time.Ticker
	tolerance uint64

	events map[gulid.ID]event.DTO
	ack    <-chan gulid.ID
	event  chan event.DTO
}

func New(c *client.C, ack <-chan gulid.ID) *R {
	return &R{
		C:       c,
		logger:  log.With().Str("app", "reader").Logger(),
		Scanner: bufio.NewScanner(os.Stdin),
		event:   make(chan event.DTO),
		events:  make(map[gulid.ID]event.DTO),
		ack:     ack,
	}
}

func (r *R) Close() error {
	if err := r.C.Close(); err != nil {
		return err
	}
	close(r.event)
	r.ticker.Stop()
	return nil
}

// Dial initialize a R.
func (r *R) Dial(cfg Config) error {
	r.tolerance = cfg.Tolerance
	var err error
	if r.addr, err = net.ResolveUDPAddr("udp", cfg.Address); err != nil {
		return err
	}

	d := time.Second / time.Duration(r.tolerance)
	r.ticker = time.NewTicker(d)
	go r.HandleACK()
	return nil
}

// Send sends message on appropriate config.
func (r R) Send(input event.DTO) {
	raw, err := input.Marshal()
	if err != nil {
		r.logger.Error().Err(err).Msg("failed to marshal action")
		return
	}
	r.event <- input
	r.logger.Info().Str("data", fmt.Sprintf("%v", input)).Msg("send")
	go r.C.Send(raw, r.addr)
}

// HandleACK handles events sending and received acks.
func (r R) HandleACK() {
	d := uint64(time.Second / time.Duration(r.tolerance))
	for {
		select {
		case <-r.ticker.C:
			now := ulid.Now()
			for _, e := range r.events {
				t := e.ID.Time()
				if t > now || now-t < d {
					continue
				}
				go func(e event.DTO) {
					raw, err := e.Marshal()
					if err != nil {
						r.logger.Error().Err(err).Msg("failed to marshal action")
						return
					}
					r.C.Send(raw, r.addr)
				}(e)
			}
		case e := <-r.event:
			r.logger.Info().Str("id", e.ID.String()).Msg("event received")
			r.events[e.ID] = e
		case id := <-r.ack:
			r.logger.Info().Str("id", id.String()).Msg("ack received")
			delete(r.events, id)
		}
	}
}
