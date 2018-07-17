package login

import (
	"net"
	"time"

	"github.com/rs/zerolog"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Renderer is the 2D graphic renderer for login/PC page.
type Renderer struct {
	logger *zerolog.Logger

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

	signinURL net.Addr
	tolerance time.Duration
	tickrate  uint32
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

// Dial initializes render window.
func (r *Renderer) Dial(cfg Config) error {
	var err error

	sdl.Do(func() {
		if err != nil {
			return
		}
		r.width, r.height = r.window.GetSize()
		r.renderer, err = sdl.CreateRenderer(r.window, -1, sdl.RENDERER_ACCELERATED)
		if err != nil {
			return
		}
		r.renderer.Clear()
		r.background, err = img.Load(cfg.Background)
		if err != nil {
			return
		}
		r.texture, err = r.renderer.CreateTextureFromSurface(r.background)
		if err != nil {
			return
		}
		if err := ttf.Init(); err != nil {
			return
		}

		if r.logFont, err = ttf.OpenFont("assets/skorzhen.ttf", 64); err != nil {
			return
		}
		if r.loginSurf, err = r.logFont.RenderUTF8Solid("login", sdl.Color{255, 255, 255, 255}); err != nil {
			return
		}
		if r.loginText, err = r.renderer.CreateTextureFromSurface(r.loginSurf); err != nil {
			return
		}
		if r.passwordSurf, err = r.logFont.RenderUTF8Solid("password", sdl.Color{255, 255, 255, 255}); err != nil {
			return
		}
		if r.passwordText, err = r.renderer.CreateTextureFromSurface(r.passwordSurf); err != nil {
			return
		}

		if r.font, err = ttf.OpenFont("assets/geosteam.ttf", 256); err != nil {
			return
		}
		if r.titleSurf, err = r.font.RenderUTF8Solid("GAME_01", sdl.Color{0, 255, 0, 255}); err != nil {
			return
		}
		if r.titleText, err = r.renderer.CreateTextureFromSurface(r.titleSurf); err != nil {
			return
		}
	})
	if err != nil {
		return err
	}
	sdl.Do(func() { go r.render() })
	return nil
}

// Close closes the render window.
func (r *Renderer) Close() error {
	sdl.Do(func() {
		if r.window != nil {
			r.window.Destroy()
		}
		if r.renderer != nil {
			r.renderer.Destroy()
		}
		if r.background != nil {
			r.background.Free()
		}
		if r.texture != nil {
			r.texture.Destroy()
		}
		if r.font != nil {
			r.font.Close()
		}
		if r.loginSurf != nil {
			r.loginSurf.Free()
		}
		if r.loginText != nil {
			r.loginText.Destroy()
		}
		if r.passwordSurf != nil {
			r.passwordSurf.Free()
		}
		if r.passwordText != nil {
			r.passwordText.Destroy()
		}
		if r.logFont != nil {
			r.logFont.Close()
		}
		if r.titleSurf != nil {
			r.titleSurf.Free()
		}
		if r.titleText != nil {
			r.titleText.Destroy()
		}
	})
	return nil
}

// render is an sdl dependant function to render current frame.
func (r *Renderer) render() {
	for {
		r.renderer.Clear()
		r.renderer.Copy(
			r.texture,
			&sdl.Rect{X: 0, Y: 0, W: 3600, H: 1800},
			&sdl.Rect{X: 0, Y: 0, W: r.width, H: r.height},
		)
		r.renderer.Copy(
			r.titleText,
			nil,
			&sdl.Rect{X: 200, Y: 42, W: 560, H: 84},
		)
		r.renderer.Copy(
			r.loginText,
			nil,
			&sdl.Rect{X: 42, Y: 220, W: 142, H: 32},
		)
		r.renderer.Copy(
			r.passwordText,
			nil,
			&sdl.Rect{X: 42, Y: 260, W: 142, H: 32},
		)
		r.renderer.Present()
		sdl.Delay(1000 / r.tickrate)
	}
}

// UnstackEvent sends and event to server.
func (r *Renderer) UnstackEvent() {
	for {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e.(type) {
			case *sdl.QuitEvent:
				return
			}
			// r.events[e.ID] = e
		}
	}
}
