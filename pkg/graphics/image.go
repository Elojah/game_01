package graphics

// Image wraps a SDL image display and memory.
type Image struct {
	Texture *sdl.Texture

	surface *sdl.Surface
}

// NewImage returns a new text object.
func NewImage(path string) (*Image, error) {
	var img Image
	var err error
	img.surface, err = img.Load(path)
	return &img, err
}

// Init initializes the Texture via renderer.
func (img *Image) Init(renderer *sdl.Renderer) error {
	var err error
	img.Texture, err = renderer.CreateTextureFromSurface(img.surface)
	return err
}

// Close destroys a text surface/Texture.
func (img *Image) Close() {
	if img.surface != nil {
		img.surface.Free()
	}
	if img.Texture != nil {
		img.Texture.Destroy()
	}
}
