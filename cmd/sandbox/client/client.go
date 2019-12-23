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
	Network *network.Client
	Token   gulid.ID

	Add    chan Sendable
	Remove chan gulid.ID
	ticker *time.Ticker
	done   chan struct{}

	sendables map[gulid.ID]Sendable
}

func NewClient() *C {
	return &C{
		Add:       make(chan Sendable),
		Remove:    make(chan gulid.ID),
		sendables: make(map[gulid.ID]Sendable),
		done:      make(chan struct{}),
	}
}

// Dial initialize a client.
func (c *C) Dial(cfg Config) error {
	c.ticker = time.NewTicker(time.Duration(cfg.Interval) * time.Millisecond)
	go c.Run()
	return nil
}

func (c *C) Close() error {
	c.done <- struct{}{}
	close(c.Add)
	close(c.Remove)
	close(c.done)
	c.ticker.Stop()
	return nil
}

func (c *C) Run() {
	for {
		select {
		case <-c.done:
			return
		case se := <-c.Add:
			c.sendables[se.ID()] = se
		case id := <-c.Remove:
			delete(c.sendables, id)
		case <-c.ticker.C:
			for _, se := range c.sendables {
				c.Network.Send(event.DTO{
					ID:    gulid.NewID(),
					Token: c.Token,
					Query: se.Query(),
				})
			}
		}
	}
}
