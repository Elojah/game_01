package renderer

import (
	"github.com/veandco/go-sdl2/sdl"
)

// R is the main 2D graphic client renderer.
type R struct {
	window *sdl.Window
}

// Dial initializes render window.
func (r *R) Dial(config Config) error {
	return nil
}
