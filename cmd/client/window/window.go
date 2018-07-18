package window

import (
	"github.com/veandco/go-sdl2/sdl"
)

// W alias a sdl window..
type W struct {
	*sdl.Window
}

// NewWindow returns a new SDL window.
func NewWindow() *W {
	return &W{}
}

// Dial initializes render window.
func (w *W) Dial(cfg Config) error {
	var err error

	sdl.Do(func() {
		w.Window, err = sdl.CreateWindow(cfg.Title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, cfg.Width, cfg.Height, sdl.WINDOW_OPENGL)
	})
	return err
}

// Close closes the render window.
func (w *W) Close() error {
	sdl.Do(func() {
		w.Window.Destroy()
	})
	return nil
}
