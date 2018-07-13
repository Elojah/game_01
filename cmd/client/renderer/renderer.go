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
	window   *sdl.Window
	renderer *sdl.Renderer

	token     ulid.ID
	addr      net.Addr
	tolerance time.Duration
	tickrate  uint32

	ticker    *time.Ticker
	ackC      <-chan infra.ACK
	entitiesC <-chan entity.E
	events    map[ulid.ID]dto.Event
	entities  map[ulid.ID]entity.E
}

// NewRenderer returns a valid renderer.
func NewRenderer(
	c *client.C,
	ackC <-chan infra.ACK,
	entitiesC <-chan entity.E,
) *R {
	return &R{
		C:         c,
		ackC:      ackC,
		entitiesC: entitiesC,
	}
}

// Dial initializes render window.
func (r *R) Dial(cfg Config) error {
	var err error

	r.token = cfg.Token
	r.tickrate = cfg.TickRate
	if r.addr, err = net.ResolveUDPAddr("udp", cfg.Address); err != nil {
		return err
	}
	r.ticker = time.NewTicker(cfg.Tolerance)
	go r.unstackEntities()
	go r.unstackACK()
	go r.resendEvent()
	sdl.Do(func() {
		r.window, err = sdl.CreateWindow(cfg.Title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, cfg.Width, cfg.Height, sdl.WINDOW_OPENGL)
		if err != nil {
			return
		}
		r.renderer, err = sdl.CreateRenderer(r.window, -1, sdl.RENDERER_ACCELERATED)
		if err != nil {
			return
		}
		r.renderer.Clear()
	})
	if err != nil {
		return err
	}
	sdl.Do(func() { go r.render() })
	return nil
}

// Close closes the render window.
func (r *R) Close() error {
	r.ticker.Stop()
	sdl.Do(func() {
		r.window.Destroy()
		r.renderer.Destroy()
	})
	if err := r.C.Close(); err != nil {
		return err
	}
	return nil
}

// UnstackEvent sends and event to server.
func (r *R) UnstackEvent() {
	for {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e.(type) {
			case *sdl.QuitEvent:
				return
			}
			// r.events[e.ID] = e
			// go func(e dto.Event) {
			// 	raw, err := e.Marshal(nil)
			// 	if err != nil {
			// 		r.logger.Error().Err(err).Msg("failed to marshal action")
			// 		return
			// 	}
			// 	r.Send(raw, r.addr)
			// }(e)
		}
	}
}

// unstackEntities read entities from chan and add them to render map.
func (r *R) unstackEntities() {
	for e := range r.entitiesC {
		r.entities[e.ID] = e
	}
}

// unstackACK read acks from chan and remove corresponding events.
func (r *R) unstackACK() {
	for ack := range r.ackC {
		delete(r.events, ack.ID)
	}
}

// resendEvent send an event again if no ack has been received since tolerance.
func (r *R) resendEvent() {
	for t := range r.ticker.C {
		for _, e := range r.events {
			if t.Sub(time.Unix(0, e.TS)) < r.tolerance {
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

// render is an sdl dependant function to render current frame.
func (r *R) render() {
	for {
		for _, e := range r.entities {
			r.renderer.SetDrawColor(0, 0, 0, 0x20)
			r.renderer.FillRect(&sdl.Rect{0, 0, int32(e.Position.Coord.X), int32(e.Position.Coord.Y)})
		}
		r.renderer.Present()
		r.renderer.Clear()
		sdl.Delay(1000 / r.tickrate)
	}
}
