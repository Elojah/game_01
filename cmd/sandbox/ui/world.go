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

	ClientSystem *ClientSystem
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

	cs := ControlSystem{}
	cs.Setup()
	w.AddSystem(&cs)
	w.AddSystem(sc.ClientSystem)

	engoc.SetBackground(color.Black)

	for _, mapname := range maps {
		ts, err := NewTiles(mapname)
		if err != nil {
			logger.Error().Err(err).Str("name", mapname).Msg("failed to load map")
			return
		}
		ts.AddToWorld(w)
	}

	e := NewEntity(chars[0])
	e.LoadAnimations()
	e.SpaceComponent.Position.X = float32(400)
	e.SpaceComponent.Position.Y = float32(400)
	e.AddToWorld(w)
}
