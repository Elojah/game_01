package main

import (
	sciter "github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

type app struct {
	w *window.Window
}

func (a *app) Dial(c Config) error {
	var err error
	a.w, err = window.New(
		sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN|sciter.SW_ENABLE_DEBUG,
		nil,
	)
	if err != nil {
		return err
	}
	if err := a.w.LoadFile(c.HTMLFile); err != nil {
		return err
	}
	a.w.SetTitle("game_launcher")
	a.w.Show()
	return nil
}

func (a *app) Close() {
}

func (a *app) Start() {
	a.w.Run()
}
