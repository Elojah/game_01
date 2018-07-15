package booter

import (
	"net"
	"time"

	"github.com/rs/zerolog"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// B is the main 2D graphic client renderer.
type B struct {
	logger zerolog.Logger

	window   *sdl.Window
	renderer *sdl.Renderer

	background *sdl.Surface
	texture    *sdl.Texture
	font       *ttf.Font

	logFont *ttf.Font

	loginSurf *sdl.Surface
	loginText *sdl.Texture

	passwordSurf *sdl.Surface
	passwordText *sdl.Texture

	titleSurf *sdl.Surface
	titleText *sdl.Texture

	width  int32
	height int32

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

	b.tickrate = cfg.TickRate
	if b.addr, err = net.ResolveTCPAddr("tcp", cfg.Address); err != nil {
		return err
	}
	sdl.Do(func() {
		sdl.Init(sdl.INIT_VIDEO)
		b.window, err = sdl.CreateWindow(cfg.Title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, cfg.Width, cfg.Height, sdl.WINDOW_OPENGL)
		if err != nil {
			return
		}
		b.width, b.height = b.window.GetSize()
		b.renderer, err = sdl.CreateRenderer(b.window, -1, sdl.RENDERER_ACCELERATED)
		if err != nil {
			return
		}
		b.renderer.Clear()
		b.background, err = img.Load(cfg.Background)
		if err != nil {
			return
		}
		b.texture, err = b.renderer.CreateTextureFromSurface(b.background)
		if err != nil {
			return
		}
		if err := ttf.Init(); err != nil {
			return
		}

		if b.logFont, err = ttf.OpenFont("assets/skorzhen.ttf", 64); err != nil {
			return
		}
		if b.loginSurf, err = b.logFont.RenderUTF8Solid("login", sdl.Color{255, 255, 255, 255}); err != nil {
			return
		}
		if b.loginText, err = b.renderer.CreateTextureFromSurface(b.loginSurf); err != nil {
			return
		}
		if b.passwordSurf, err = b.logFont.RenderUTF8Solid("password", sdl.Color{255, 255, 255, 255}); err != nil {
			return
		}
		if b.passwordText, err = b.renderer.CreateTextureFromSurface(b.passwordSurf); err != nil {
			return
		}

		if b.font, err = ttf.OpenFont("assets/geosteam.ttf", 256); err != nil {
			return
		}
		if b.titleSurf, err = b.font.RenderUTF8Solid("GAME_01", sdl.Color{0, 255, 0, 255}); err != nil {
			return
		}
		if b.titleText, err = b.renderer.CreateTextureFromSurface(b.titleSurf); err != nil {
			return
		}
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
		if b.window != nil {
			b.window.Destroy()
		}
		if b.renderer != nil {
			b.renderer.Destroy()
		}
		if b.background != nil {
			b.background.Free()
		}
		if b.texture != nil {
			b.texture.Destroy()
		}
		if b.font != nil {
			b.font.Close()
		}
		if b.loginSurf != nil {
			b.loginSurf.Free()
		}
		if b.loginText != nil {
			b.loginText.Destroy()
		}
		if b.passwordSurf != nil {
			b.passwordSurf.Free()
		}
		if b.passwordText != nil {
			b.passwordText.Destroy()
		}
		if b.logFont != nil {
			b.logFont.Close()
		}
		if b.titleSurf != nil {
			b.titleSurf.Free()
		}
		if b.titleText != nil {
			b.titleText.Destroy()
		}
	})
	return nil
}

// render is an sdl dependant function to render current frame.
func (b *B) render() {
	for {
		b.renderer.Clear()
		b.renderer.Copy(
			b.texture,
			&sdl.Rect{X: 0, Y: 0, W: 3600, H: 1800},
			&sdl.Rect{X: 0, Y: 0, W: b.width, H: b.height},
		)
		b.renderer.Copy(
			b.titleText,
			nil,
			&sdl.Rect{X: 200, Y: 42, W: 560, H: 84},
		)
		b.renderer.Copy(
			b.loginText,
			nil,
			&sdl.Rect{X: 42, Y: 220, W: 142, H: 32},
		)
		b.renderer.Copy(
			b.passwordText,
			nil,
			&sdl.Rect{X: 42, Y: 260, W: 142, H: 32},
		)
		b.renderer.Present()
		sdl.Delay(1000 / b.tickrate)
	}
}

// UnstackEvent sends and event to server.
func (b *B) UnstackEvent() {
	for {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e.(type) {
			case *sdl.QuitEvent:
				return
			}
			// b.events[e.ID] = e
		}
	}
}
