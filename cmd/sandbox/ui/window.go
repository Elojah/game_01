package ui

import (
	"image"
	"io/ioutil"
	"path"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	frameW = 125
	frameH = 200

	frameN       = 40
	frameLines   = 5
	frameColumns = 8
)

var (
	currentFrame = 0
)

// Window is the UI container.
type Window struct {
	Width  int
	Height int

	AssetsDir string

	images map[string]*ebiten.Image
}

// LoadAssets loads all assets present in assets directory.
func (w Window) LoadAssets() error {

	files, err := ioutil.ReadDir(path.Join(w.AssetsDir, "img"))
	if err != nil {
		return err
	}

	w.images = make(map[string]*ebiten.Image)

	for _, f := range files {
		name := f.Name()
		ei, _, err := ebitenutil.NewImageFromFile(path.Join(w.AssetsDir, "img", name), ebiten.FilterDefault)
		if err != nil {
			return err
		}
		w.images[name] = ei
	}

	return nil
}

func (w Window) Up() error {
	return ebiten.Run(w.Update, w.Width, w.Height, 1, "GAME_01")
}

func (w Window) Update(screen *ebiten.Image) error {

	if currentFrame >= 40 {
		currentFrame = 0
	}

	if ebiten.IsDrawingSkipped() {
		currentFrame++
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameW)/2, -float64(frameH)/2)
	op.GeoM.Translate(float64(w.Width)/2, float64(w.Height)/2)
	// i := currentFrame % frameColumns
	// j := currentFrame / frameColumns
	// sx, sy := i*frameW, j*frameH
	// screen.DrawImage(w.images["midnight.png"].SubImage(image.Rect(sx, sy, sx+frameW, sy+frameH)).(*ebiten.Image), op)

	screen.DrawImage(w.images["midnight.png"].SubImage(image.Rect(0, 0, frameW, frameH)).(*ebiten.Image), op)

	currentFrame++
	return nil
}
