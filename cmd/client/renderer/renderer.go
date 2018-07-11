package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Renderer is the main 2D graphic client renderer.
type Renderer struct {
	window *sdl.Window
}

// Dial initializes render window.
func (r *Renderer) Dial(config Config) error {
	return nil
}
