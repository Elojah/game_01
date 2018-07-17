package login

import (
	"net"
	"time"

	"github.com/elojah/game_01/pkg/graphics"
	"github.com/rs/zerolog"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Renderer is the 2D graphic renderer for login/PC page.
type Renderer struct {
	logger *zerolog.Logger

	renderer *sdl.Renderer

	backgroundImg *graphics.Image

	skorzhenFont *ttf.Font
	loginText    *graphics.Text
	passwordText *graphics.Text

	geosteamFont *ttf.Font
	titleText    *graphics.Text

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
		if r.backgroundImg, err = graphics.NewImage(cfg.Background); err != nil {
			return
		}

		// Init fonts
		if err := ttf.Init(); err != nil {
			return
		}
		if r.skorzhenFont, err = ttf.OpenFont("assets/skorzhen.ttf", 64); err != nil {
			return
		}
		if r.geosteamFont, err = ttf.OpenFont("assets/geosteam.ttf", 64); err != nil {
			return
		}

		// Init texts
		if r.loginText, err = graphics.NewText("login", skorzhenFont, sdl.Color{44, 56, 126}); err != nil {
			return
		}
		if r.passwordText, err = graphics.NewText("password", skorzhenFont, sdl.Color{44, 56, 126}); err != nil {
			return
		}
		if r.titleText, err = graphics.NewText("GAME_01", geosteamFont, sdl.Color{178, 42, 0}); err != nil {
			return
		}
	})
	return err
}

// Close closes the render window.
func (r *Renderer) Close() error {
	sdl.Do(func() {
		if r.renderer != nil {
			r.renderer.Destroy()
		}
		if r.skorzhenFont != nil {
			r.skorzhenFont.Close()
		}
		if r.geosteamFont != nil {
			r.geosteamFont.Close()
		}
		if r.loginText != nil {
			r.loginText.Close()
		}
		if r.passwordText != nil {
			r.passwordText.Close()
		}
		if r.titleText != nil {
			r.titleText.Close()
		}
	})
	return nil
}

func (r *Renderer) Init(w *sdl.Window) error {
	var err error
	sdl.Do(func() {
		if r.renderer, err = sdl.CreateRenderer(w, -1, sdl.RENDERER_ACCELERATED); err != nil {
			return
		}
		if err := r.backgroundImg.Init(r.renderer); err != nil {
			return
		}
		if err := r.loginText.Init(r.renderer); err != nil {
			return
		}
		if err := r.passwordText.Init(r.renderer); err != nil {
			return
		}
		if err := r.titleText.Init(r.renderer); err != nil {
			return
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

// // UnstackEvent sends and event to server.
// func (r *Renderer) UnstackEvent() {
// 	for {
// 		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
// 			switch e.(type) {
// 			case *sdl.QuitEvent:
// 				return
// 			}
// 			// r.events[e.ID] = e
// 		}
// 	}
// }
