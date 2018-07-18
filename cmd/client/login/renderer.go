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

	width  int32
	height int32

	backgroundImg *graphics.Image

	bitwiseFont  *ttf.Font
	loginText    *graphics.Text
	passwordText *graphics.Text

	geosteamFont *ttf.Font
	titleText    *graphics.Text

	mozart0Font   *ttf.Font
	loginInput    *graphics.TextInput
	passwordInput *graphics.TextInput
	focus         *graphics.TextInput

	signinURL net.Addr
	tolerance time.Duration
	tickrate  uint32
}

// NewRenderer returns a new login renderer.
func NewRenderer() *Renderer {
	return &Renderer{}
}

// Dial initializes render window.
func (r *Renderer) Dial(cfg Config) error {
	var err error

	sdl.Do(func() {
		// Init images
		if r.backgroundImg, err = graphics.NewImage(cfg.Background); err != nil {
			return
		}

		// Init fonts
		if err := ttf.Init(); err != nil {
			return
		}
		if r.bitwiseFont, err = ttf.OpenFont("assets/bitwise.ttf", 64); err != nil {
			return
		}
		if r.geosteamFont, err = ttf.OpenFont("assets/geosteam.ttf", 256); err != nil {
			return
		}
		if r.mozart0Font, err = ttf.OpenFont("assets/mozart_0.ttf", 64); err != nil {
			return
		}

		// Init texts
		if r.loginText, err = graphics.NewText("login", r.bitwiseFont, sdl.Color{57, 255, 20, 255}); err != nil {
			return
		}
		if r.passwordText, err = graphics.NewText("password", r.bitwiseFont, sdl.Color{57, 255, 20, 255}); err != nil {
			return
		}
		if r.titleText, err = graphics.NewText("GAME_01", r.geosteamFont, sdl.Color{178, 42, 0, 255}); err != nil {
			return
		}

		// Text inputs
		r.loginInput = graphics.NewTextInput(r.mozart0Font, sdl.Color{57, 255, 20, 255})
		r.passwordInput = graphics.NewTextInput(r.mozart0Font, sdl.Color{57, 255, 20, 255})
		r.focus = r.loginInput
	})
	return err
}

// Close closes the render window.
func (r *Renderer) Close() error {
	sdl.Do(func() {
		if r.renderer != nil {
			r.renderer.Destroy()
		}
		if r.bitwiseFont != nil {
			r.bitwiseFont.Close()
		}
		if r.geosteamFont != nil {
			r.geosteamFont.Close()
		}
		if r.mozart0Font != nil {
			r.mozart0Font.Close()
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
		if r.loginInput != nil {
			r.loginInput.Close()
		}
		if r.passwordInput != nil {
			r.passwordInput.Close()
		}
	})
	return nil
}

// Init initializes the login page renderer.
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
		r.width, r.height = w.GetSize()
	})
	return nil
}

// Update is an sdl dependant function to render current frame.
func (r *Renderer) Update() {
	sdl.StartTextInput()
	for {
		r.renderer.Clear()
		r.renderer.Copy(
			r.backgroundImg.Texture,
			&sdl.Rect{X: 0, Y: 0, W: 3600, H: 1800},
			&sdl.Rect{X: 0, Y: 0, W: r.width, H: r.height},
		)
		r.renderer.Copy(
			r.titleText.Texture,
			nil,
			&sdl.Rect{X: r.width / 5, Y: r.height / 10, W: 3 * r.width / 5, H: r.height / 5},
		)
		r.renderer.Copy(
			r.loginText.Texture,
			nil,
			&sdl.Rect{X: r.width / 10, Y: r.height/2 - r.height/15, W: r.width / 10, H: r.height / 15},
		)
		r.renderer.Copy(
			r.passwordText.Texture,
			nil,
			&sdl.Rect{X: r.width / 10, Y: r.height / 2, W: r.width / 5, H: r.height / 15},
		)
		r.loginInput.Update(r.renderer)
		r.passwordInput.Update(r.renderer)
		r.renderer.Copy(
			r.loginInput.Texture,
			nil,
			&sdl.Rect{X: r.width / 3, Y: r.height/2 - r.height/15, W: r.width / 2, H: r.height / 15},
		)
		r.renderer.Copy(
			r.passwordInput.Texture,
			nil,
			&sdl.Rect{X: r.width / 3, Y: r.height / 2, W: r.width / 2, H: r.height / 15},
		)
		r.renderer.Present()
		sdl.Delay(180)
	}
}

// PollEvent sends and event to server.
func (r *Renderer) PollEvent() {
	for {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.KeyboardEvent:
				r.focus.Input(e.(*sdl.KeyboardEvent))
			}
			// r.events[e.ID] = e
		}
	}
}
