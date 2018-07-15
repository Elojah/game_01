package booter

import (
	"net"
	"time"

	"github.com/rs/zerolog"
	"github.com/veandco/go-sdl2/sdl"
)

// B is the main 2D graphic client renderer.
type B struct {
	logger zerolog.Logger

	window   *sdl.Window
	renderer *sdl.Renderer

	addr      net.Addr
	tolerance time.Duration
	tickrate  uint32
}

// NewBooter returns a valid renderer.
func NewBooter() *B {
	return &B{}
}

// Dial initializes render window.
func (b *B) Dial(cfg Config) error {
	var err error

	if b.addr, err = net.ResolveTCPAddr("tcp", cfg.Address); err != nil {
		return err
	}
	sdl.Do(func() {
		b.window, err = sdl.CreateWindow(cfg.Title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, cfg.Width, cfg.Height, sdl.WINDOW_OPENGL)
		if err != nil {
			return
		}
		b.renderer, err = sdl.CreateRenderer(b.window, -1, sdl.RENDERER_ACCELERATED)
		if err != nil {
			return
		}
		b.renderer.Clear()
	})
	if err != nil {
		return err
	}
	sdl.Do(func() { go b.render() })
	return nil
}

// Close closes the render window.
func (b *B) Close() error {
	sdl.Do(func() {
		b.window.Destroy()
		b.renderer.Destroy()
	})
	return nil
}

// render is an sdl dependant function to render current frame.
func (b *B) render() {
	for {
		b.renderer.Present()
		sdl.Delay(1000 / b.tickrate)
	}
}
