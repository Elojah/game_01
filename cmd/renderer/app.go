package main

import (
	"time"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type app struct {
	config Config

	win    *pixelgl.Window
	ticker *time.Ticker
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
		}
		a.win.Update()
	}
}

func (a *app) Close() {
	a.ticker.Stop()
}
