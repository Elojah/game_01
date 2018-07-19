package graphics

import (
	"strings"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// TextInput represents a flat text input field
type TextInput struct {
	Texture *sdl.Texture

	Previous *TextInput
	Next     *TextInput

	surface *sdl.Surface

	content []rune
	color   sdl.Color
	font    *ttf.Font
	max     uint
	Hidden  bool

	cursor     uint
	showCursor bool
}

// NewTextInput returns a new text object.
func NewTextInput(font *ttf.Font, color sdl.Color, max uint) *TextInput {
	return &TextInput{
		font:    font,
		color:   color,
		max:     max,
		content: make([]rune, max+1),
	}
}

// Update updates the Texture via renderer.
func (t *TextInput) Update(renderer *sdl.Renderer) error {
	var err error
	t.Close()
	if t.surface, err = t.font.RenderUTF8Solid(t.RightPad(), t.color); err != nil {
		return err
	}
	if t.Texture, err = renderer.CreateTextureFromSurface(t.surface); err != nil {
		return err
	}
	// if t.showCursor {
	// 	if (time.Now().UnixNano() % 1000000000) < 500000000 {
	// 		t.content[t.cursor] = '_'
	// 	} else {
	// 		t.content[t.cursor] = ' '
	// 	}
	// }
	return nil
}

// RightPad returns the actual content right padded with spaces to max.
func (t *TextInput) RightPad() string {
	if t.Hidden {
		return strings.Repeat("•", int(t.cursor)) + string(t.content[t.cursor]) + strings.Repeat(" ", int(t.max-t.cursor))
	}
	return string(t.content[:t.cursor+1]) + strings.Repeat(" ", int(t.max-t.cursor))
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
		t.content[t.cursor] = ' '
		t.content[t.cursor-1] = ' '
		t.cursor--
		return nil
	}
	r, ok := keymap[input.Keysym.Sym]
	if !ok {
		return nil
	}
	if t.cursor >= t.max {
		return nil
	}
	t.content[t.cursor] = r
	t.cursor++
	return nil
}

// ShowCursor set if the text input show a text cursor.
func (t *TextInput) ShowCursor(show bool) {
	t.content[t.cursor] = ' '
	t.showCursor = show
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
