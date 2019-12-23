package network

import (
	"bufio"
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

type Client struct {
	*client.C

	logger zerolog.Logger
	*bufio.Scanner

	addr net.Addr

	ticker        *time.Ticker
	ackTolerance  uint64
	omitTolerance uint64

	events map[gulid.ID]event.DTO
	ack    <-chan gulid.ID
	event  chan event.DTO
}

func New(c *client.C, ack <-chan gulid.ID) *Client {
	return &Client{
		C:       c,
		logger:  log.With().Str("app", "network").Logger(),
		Scanner: bufio.NewScanner(os.Stdin),
		event:   make(chan event.DTO),
		events:  make(map[gulid.ID]event.DTO),
		ack:     ack,
	}
}

func (c *Client) Close() error {
	if err := c.C.Close(); err != nil {
		return err
	}
	close(c.event)
	c.ticker.Stop()
	return nil
}

// Dial initialize a Client.
func (c *Client) Dial(cfg Config) error {
	c.ackTolerance = cfg.ACKTolerance
	c.omitTolerance = cfg.OmitTolerance
	var err error
	if c.addr, err = net.ResolveUDPAddr("udp", cfg.Address); err != nil {
		return err
	}

	c.ticker = time.NewTicker(time.Duration(int(c.ackTolerance)) * time.Millisecond)
	go c.HandleACK()
	return nil
}

// Send sends message on appropriate config.
func (c Client) Send(input event.DTO) {
	raw, err := input.Marshal()
	if err != nil {
		c.logger.Error().Err(err).Msg("failed to marshal action")
		return
	}
	c.event <- input
	go c.C.Send(raw, c.addr)
}

// HandleACK handles events sending and received acks.
func (c Client) HandleACK() {
	for {
		select {
		case id := <-c.ack:
			c.logger.Info().Str("id", id.String()).Msg("event acked")
			delete(c.events, id)
		case e := <-c.event:
			c.logger.Info().Str("id", e.ID.String()).Msg("event registered")
			c.events[e.ID] = e
		case <-c.ticker.C:
			now := ulid.Now()
			for _, e := range c.events {
				t := e.ID.Time()

				if t > now {
					// if event time is in the future, wait for ack
					continue
				} else if now-t < c.ackTolerance {
					// if event time is within ack tolerance, wait for ack
					continue
				} else if now-t > c.omitTolerance {
					// if event time is after omit tolerance, discard it
					c.logger.Info().Str("id", e.ID.String()).Msg("event omitted")
					delete(c.events, e.ID)
					continue
				}

				raw, err := e.Marshal()
				if err != nil {
					c.logger.Error().Err(err).Msg("failed to marshal action")
					continue
				}
				c.C.Send(raw, c.addr)
			}
		}
	}
}
