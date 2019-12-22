package client

import (
	"time"

	"github.com/elojah/game_01/cmd/sandbox/reader"
	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

type Queryable interface {
	ID() gulid.ID
	Query() event.Query
}

type Client struct {
	Add    chan Queryable
	Remove chan gulid.ID
	ticker *time.Ticker
	close  chan struct{}

	Reader *reader.R

	queries map[gulid.ID]Queryable
}

func NewClient() *Client {
	return &Client{
		Add:    make(chan Queryable),
		Remove: make(chan gulid.ID),

		queries: make(map[gulid.ID]Queryable),
	}
}

// Dial initialize a client.
func (c *Client) Dial(cfg Config) error {
	c.ticker = time.NewTicker(time.Duration(cfg.Interval) * time.Millisecond)
	go c.Run()
	return nil
}

func (c *Client) Close() error {
	close(c.Add)
	close(c.Remove)
	c.close <- struct{}{}
	c.ticker.Stop()
	return nil
}

func (c *Client) Run() {
	for {
		select {
		case <-c.close:
			return
		case <-c.ticker.C:
			for _, q := range c.queries {
				c.Reader.Send(event.DTO{
					ID:    gulid.NewID(),
					Token: gulid.NewID(),
					Query: q.Query(),
				})
			}
		case q := <-c.Add:
			c.queries[q.ID()] = q
		case id := <-c.Remove:
			delete(c.queries, id)
		}
	}
}
