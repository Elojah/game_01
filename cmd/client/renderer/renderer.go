package renderer

import (
	"net"
	"time"

	"github.com/rs/zerolog"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/elojah/game_01/pkg/dto"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/mux/client"
)

// R is the main 2D graphic client renderer.
type R struct {
	logger zerolog.Logger

	*client.C
	window *sdl.Window

	token     ulid.ID
	addr      net.Addr
	tolerance time.Duration

	ticker    *time.Ticker
	inputC    <-chan dto.Event
	ackC      <-chan infra.ACK
	entitiesC <-chan entity.E
	inputs    map[ulid.ID]dto.Event
	entities  map[ulid.ID]entity.E
}

// NewRenderer returns a valid renderer.
func NewRenderer(
	c *client.C,
	inputC <-chan dto.Event,
	ackC <-chan infra.ACK,
	entitiesC <-chan entity.E,
) *R {
	return &R{
		C:         c,
		inputC:    inputC,
		ackC:      ackC,
		entitiesC: entitiesC,
	}
}

// Dial initializes render window.
func (r *R) Dial(cfg Config) error {
	var err error

	r.token = cfg.Token
	if r.addr, err = net.ResolveUDPAddr("udp", cfg.Address); err != nil {
		return err
	}
	r.ticker = time.NewTicker(r.tolerance)
	go r.UnstackEntities()
	go r.UnstackACK()
	go r.ResendEvent()
	return nil
}

// Close closes the render window.
func (r *R) Close() error {
	r.ticker.Stop()
	if err := r.C.Close(); err != nil {
		return err
	}
	return nil
}

// UnstackEntities read entities from chan and add them to render map.
func (r *R) UnstackEntities() {
	for e := range r.entitiesC {
		r.entities[e.ID] = e
	}
}

// UnstackACK read acks from chan and remove corresponding inputs.
func (r *R) UnstackACK() {
	for ack := range r.ackC {
		delete(r.inputs, ack.ID)
	}
}

// SendEvent sends and event to server.
func (r *R) SendEvent() {
	for e := range r.inputC {
		go func(e dto.Event) {
			raw, err := e.Marshal(nil)
			if err != nil {
				r.logger.Error().Err(err).Msg("failed to marshal action")
				return
			}
			r.Send(raw, r.addr)
		}(e)
	}
}

// ResendEvent send an event again if no ack has been received since tolerance.
func (r *R) ResendEvent() {
	for t := range r.ticker.C {
		now := time.Now()
		for _, e := range r.inputs {
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
	}
}
