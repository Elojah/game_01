package main

import (
	"time"

	"golang.org/x/image/colornames"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type app struct {
	config Config

	win     *pixelgl.Window
	ticker  *time.Ticker
	entityC <-chan entity.E
}

func (a *app) Dial(c Config) error {
	a.config = c
	a.ticker = time.NewTicker(time.Second / time.Duration(c.TickRate))
	return nil
}

func (a *app) Start() {
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
	a.win.Clear(colornames.Black)
	for !a.win.Closed() {
		select {
		case <-a.ticker.C:
		case e := <-a.entityC:
			imd := imdraw.New(nil)
			imd.Color = pixel.RGB(255, 0, 0)
			imd.Push(pixel.V(e.Position.Coord.X, e.Position.Coord.Y))
			imd.Circle(40, 0)
		}
		a.win.Update()
	}
}

func (a *app) Close() {
	a.ticker.Stop()
}
