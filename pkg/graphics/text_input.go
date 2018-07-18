package graphics

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// TextInput represents a flat text input field
type TextInput struct {
	Texture *sdl.Texture

	surface *sdl.Surface

	content string
	color   sdl.Color
	font    *ttf.Font

	cursor uint
}

// NewTextInput returns a new text object.
func NewTextInput(font *ttf.Font, color sdl.Color) *TextInput {
	return &TextInput{
		font:  font,
		color: color,
	}
}

// Update updates the Texture via renderer.
func (t *TextInput) Update(renderer *sdl.Renderer) error {
	var err error
	t.Close()
	if t.surface, err = t.font.RenderUTF8Solid(t.content, t.color); err != nil {
		return err
	}
	if t.Texture, err = renderer.CreateTextureFromSurface(t.surface); err != nil {
		return err
	}
	return nil
}

// Input receives a text input and add it to current content.
func (t *TextInput) Input(input *sdl.KeyboardEvent) error {
	switch input.GetType() {
	case sdl.KEYDOWN:
		t.content += sdl.GetKeyName(input.Keysym.Sym)
	}
	return nil
}

// Close destroys a text surface/Texture.
func (t *TextInput) Close() error {
	if t.surface != nil {
		t.surface.Free()
	}
	if t.Texture != nil {
		t.Texture.Destroy()
	}
	return nil
}
