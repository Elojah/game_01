package graphics

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Text represents a SDL text to attach to a renderer.
type Text struct {
	Texture *sdl.Texture

	surface *sdl.Surface
}

// NewText returns a new text object.
func NewText(content string, font *ttf.Font, color sdl.Color) (*Text, error) {
	var t Text
	var err error
	t.surface, err = font.RenderUTF8Solid(content, color)
	return &t, err
}

// Init initializes the Texture via renderer.
func (t *Text) Init(renderer *sdl.Renderer) error {
	var err error
	t.Texture, err = renderer.CreateTextureFromSurface(t.surface)
	return err
}

// Close destroys a text surface/Texture.
func (t *Text) Close() {
	if t.surface != nil {
		t.surface.Free()
	}
	if t.Texture != nil {
		t.Texture.Destroy()
	}
}
