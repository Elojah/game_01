package ui

import (
	"image/color"
	"io/ioutil"

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
	var assets []string // nolint: prealloc
	files, err := ioutil.ReadDir(sc.Assets)
	if err != nil {
		logger.Error().Err(err).Str("dir", sc.Assets).Msg("failed to read directory")
		return
	}
	engo.Files.SetRoot(sc.Assets)
	for _, f := range files {
		assets = append(assets, f.Name())
	}
	if err := engo.Files.Load(assets...); err != nil {
		logger.Error().Err(err).Msg("failed to load assets")
		return
	}
	logger.Info().Int("assets", len(assets)).Msg("assets loaded")
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (sc *Scene) Setup(u engo.Updater) {
	logger := log.With().Str("scene", "setup").Logger()

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
