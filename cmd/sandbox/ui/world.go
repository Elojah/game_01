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

	// Load FX assets
	var fxs []string // nolint: prealloc
	fxdir := path.Join(sc.Assets, "fx")
	files, err := ioutil.ReadDir(fxdir)
	if err != nil {
		logger.Error().Err(err).Str("dir", fxdir).Msg("failed to read directory")
		return
	}
	engo.Files.SetRoot(path.Join(sc.Assets, "fx"))
	for _, f := range files {
		fxs = append(fxs, f.Name())
	}
	engo.Files.Load(fxs...)
	logger.Info().Int("fxs", len(fxs)).Msg("fxs loaded")

	// Load tilemap
	var maps []string // nolint: prealloc
	mapdir := path.Join(sc.Assets, "map")
	files, err = ioutil.ReadDir(mapdir)
	if err != nil {
		logger.Error().Err(err).Str("dir", fxdir).Msg("failed to read directory")
		return
	}
	engo.Files.SetRoot(path.Join(sc.Assets, "map"))
	for _, f := range files {
		maps = append(maps, f.Name())
	}
	engo.Files.Load(maps...)
	logger.Info().Int("maps", len(maps)).Msg("maps loaded")
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (sc *Scene) Setup(u engo.Updater) {
	logger := log.With().Str("step", "setup").Logger()

	w, _ := u.(*ecs.World)

	w.AddSystem(&engoc.RenderSystem{})
	w.AddSystem(&engoc.AnimationSystem{})

	engoc.SetBackground(color.Black)

	for _, mapname := range maps {
		ts, err := NewTiles(mapname)
		if err != nil {
			logger.Error().Err(err).Msg("failed to load map")
			return
		}
		ts.AddToWorld(w)
	}

	// Show all fxs on screen
	for i, fx := range fxs {
		e := NewEntity(fx)
		e.LoadAnimations()
		j := i / 8
		i = i % 8
		e.SpaceComponent.Position.X = float32(i * 100)
		e.SpaceComponent.Position.Y = float32((j * 100) + 20)
		e.AddToWorld(w)
	}
}
