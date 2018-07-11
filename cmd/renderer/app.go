package main

import (
	"time"

	"golang.org/x/image/colornames"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type app struct {
	config Config

	entities map[ulid.ID]Entity

	win    *pixelgl.Window
	ticker *time.Ticker
}

func (a *app) Dial(c Config) error {
	a.config = c
	a.ticker = time.NewTicker(time.Second / time.Duration(c.TickRate))
	a.entities = make(map[ulid.ID]Entity)
	return nil
}

func (a *app) Start(entityC <-chan entity.E) {
	var err error
	a.win, err = pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  a.config.Title,
		Bounds: pixel.R(0, 0, a.config.Width, a.config.Height),
		VSync:  a.config.VSync,
		// TODO use Monitor for full screen
		Monitor:     nil,
		Resizable:   a.config.Resizable,
		Undecorated: a.config.Undecorated,
	})
	if err != nil {
		return
	}
	imd := imdraw.New(nil)
	for !a.win.Closed() {
		a.win.Clear(colornames.Black)
		select {
		case <-a.ticker.C:
			for _, es := range a.entities {
				es.Draw(imd)
			}
			imd.Draw(a.win)
			imd.Clear()
		case e := <-entityC:
			a.entities[e.ID] = NewEntity(e)
		}
		a.win.Update()
	}
}

func (a *app) Close() {
	a.ticker.Stop()
}
