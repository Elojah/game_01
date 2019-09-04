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

	sc.animations = make(map[string]*engoc.Animation)
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

	ent := ecs.NewBasic()
	scomp := engoc.SpaceComponent{
		Position: engo.Point{},
		Width:    150,
		Height:   150,
	}

	spsh := engoc.NewSpritesheetFromFile("phantom.png", 100, 100)
	rcomp := engoc.RenderComponent{
		Drawable: spsh.Cell(0),
		Scale:    engo.Point{3, 3},
	}
	acomp := engoc.NewAnimationComponent(spsh.Drawables(), 0.1)

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *engoc.RenderSystem:
			sys.Add(&ent, &rcomp, &scomp)
		case *engoc.AnimationSystem:
			sys.Add(&ent, &acomp, &rcomp)
			// case *ControlSystem:
			// 	sys.Add(&ent.BasicEntity, &ent.AnimationComponent)
		}
	}
}
