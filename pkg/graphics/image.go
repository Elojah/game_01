package graphics

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// Image wraps a SDL image display and memory.
type Image struct {
	Texture *sdl.Texture

	surface *sdl.Surface
}

// NewImage returns a new text object.
func NewImage(path string) (*Image, error) {
	var im Image
	var err error
	im.surface, err = img.Load(path)
	return &im, err
}

// Init initializes the Texture via renderer.
func (im *Image) Init(renderer *sdl.Renderer) error {
	var err error
	im.Texture, err = renderer.CreateTextureFromSurface(im.surface)
	return err
}

// Close destroys a text surface/Texture.
func (im *Image) Close() {
	if im.surface != nil {
		im.surface.Free()
	}
	if im.Texture != nil {
		im.Texture.Destroy()
	}
}
