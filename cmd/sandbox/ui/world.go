package ui

import (
	"image/color"
	"io/ioutil"
	"path"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	engoc "github.com/EngoEngine/engo/common"
	"github.com/rs/zerolog/log"
)

// Scene alias a engo scene.
type Scene struct {
	Assets string
}

// Type uniquely defines your game type
func (*Scene) Type() string { return "world" }

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (sc *Scene) Preload() {

	logger := log.With().Str("scene", "preload").Logger()

	// # Load fxs
	fxdir := path.Join(sc.Assets, "fx")
	files, err := ioutil.ReadDir(fxdir)
	if err != nil {
		logger.Error().Err(err).Str("dir", fxdir).Msg("failed to read directory")
		return
	}

	engo.Files.SetRoot(path.Join(sc.Assets, "fx"))
	for _, f := range files {
		name := f.Name()
		engo.Files.Load(name)
		logger.Info().Str("file", name).Msg("fx loaded")
	}

}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (sc *Scene) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)

	w.AddSystem(&engoc.RenderSystem{})
	w.AddSystem(&engoc.AnimationSystem{})

	engoc.SetBackground(color.Black)

	e := NewEntity(entities[0])
	e.LoadAnimations()
	e.AddToWorld(w)
}
