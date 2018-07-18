package graphics

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// TextInput represents a flat text input field
type TextInput struct {
	Texture *sdl.Texture

	surface *sdl.Surface

	content [32]rune
	color   sdl.Color
	font    *ttf.Font

	cursor uint
}

// NewTextInput returns a new text object.
func NewTextInput(font *ttf.Font, color sdl.Color) *TextInput {
	return &TextInput{
		font:    font,
		color:   color,
		content: [32]rune{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	}
}

// Update updates the Texture via renderer.
func (t *TextInput) Update(renderer *sdl.Renderer) error {
	var err error
	t.Close()
	if t.surface, err = t.font.RenderUTF8Solid(string(t.content[:]), t.color); err != nil {
		return err
	}
	if t.Texture, err = renderer.CreateTextureFromSurface(t.surface); err != nil {
		return err
	}
	return nil
}

// Input receives a text input and add it to current content.
func (t *TextInput) Input(input *sdl.KeyboardEvent) error {
	if input.GetType() == sdl.KEYUP {
		return nil
	}
	if _, ok := backspaces[input.Keysym.Sym]; ok {
		if t.cursor == 0 {
			return nil
		}
		t.content[t.cursor-1] = ' '
		t.cursor--
		return nil
	}
	r, ok := keymap[input.Keysym.Sym]
	if !ok {
		return nil
	}
	if t.cursor >= 32 {
		return nil
	}
	t.content[t.cursor] = r
	t.cursor++
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
