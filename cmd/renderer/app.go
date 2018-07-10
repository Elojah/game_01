package main

import (
	"time"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type app struct {
	win    *pixelgl.Window
	ticker *time.Ticker
}

func (a *app) Dial(c Config) error {
	var err error
	a.win, err = pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  c.Title,
		Bounds: pixel.R(0, 0, c.Width, c.Height),
		VSync:  c.VSync,
		// TODO use Monitor for full screen
		Monitor:     nil,
		Resizable:   c.Resizable,
		Undecorated: c.Undecorated,
	})
	if err != nil {
		return err
	}
	a.ticker = time.NewTicker(time.Second / time.Duration(c.TickRate))
	a.win.Clear(colornames.White)
	pixelgl.Run(a.Start)
	return nil
}

func (a *app) Start() {
	for !a.win.Closed() {
		select {
		case <-a.ticker.C:
		}
		a.win.Update()
	}
}

func (a *app) Close() {
	a.ticker.Stop()
}
