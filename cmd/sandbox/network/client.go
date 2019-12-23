package network

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

type Client struct {
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

func New(c *client.C, ack <-chan gulid.ID) *Client {
	return &Client{
		C:       c,
		logger:  log.With().Str("app", "reader").Logger(),
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
	c.tolerance = cfg.Tolerance
	var err error
	if c.addr, err = net.ResolveUDPAddr("udp", cfg.Address); err != nil {
		return err
	}

	d := time.Second / time.Duration(c.tolerance)
	c.ticker = time.NewTicker(d)
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
	c.logger.Info().Str("data", fmt.Sprintf("%v", input)).Msg("send")
	go c.C.Send(raw, c.addr)
}

// HandleACK handles events sending and received acks.
func (c Client) HandleACK() {
	d := uint64(time.Second / time.Duration(c.tolerance))
	for {
		select {
		case <-c.ticker.C:
			now := ulid.Now()
			for _, e := range c.events {
				t := e.ID.Time()
				if t > now || now-t < d {
					continue
				}
				go func(e event.DTO) {
					raw, err := e.Marshal()
					if err != nil {
						c.logger.Error().Err(err).Msg("failed to marshal action")
						return
					}
					c.C.Send(raw, c.addr)
				}(e)
			}
		case e := <-c.event:
			c.logger.Info().Str("id", e.ID.String()).Msg("event received")
			c.events[e.ID] = e
		case id := <-c.ack:
			c.logger.Info().Str("id", id.String()).Msg("ack received")
			delete(c.events, id)
		}
	}
}
