package client

import (
	"time"

	"github.com/elojah/game_01/cmd/sandbox/network"
	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

type Sendable interface {
	ID() gulid.ID
	Query() event.Query
}

type C struct {
	Add    chan Sendable
	Remove chan gulid.ID
	ticker *time.Ticker
	close  chan struct{}

	network *network.Client

	sendables map[gulid.ID]Sendable
}

func NewClient(nc *network.Client) *C {
	return &C{
		Add:       make(chan Sendable),
		Remove:    make(chan gulid.ID),
		sendables: make(map[gulid.ID]Sendable),

		network: nc,
	}
}

// Dial initialize a client.
func (c *C) Dial(cfg Config) error {
	c.ticker = time.NewTicker(time.Duration(cfg.Interval) * time.Millisecond)
	go c.Run()
	return nil
}

func (c *C) Close() error {
	close(c.Add)
	close(c.Remove)
	c.close <- struct{}{}
	c.ticker.Stop()
	return nil
}

func (c *C) Run() {
	for {
		select {
		case <-c.close:
			return
		case se := <-c.Add:
			c.sendables[se.ID()] = se
		case id := <-c.Remove:
			delete(c.sendables, id)
		case <-c.ticker.C:
			for _, se := range c.sendables {
				c.network.Send(event.DTO{
					ID:    gulid.NewID(),
					Token: gulid.NewID(),
					Query: se.Query(),
				})
			}
		}
	}
}
