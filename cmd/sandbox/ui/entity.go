package ui

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	engoc "github.com/EngoEngine/engo/common"
)

var (
	entities = []Entity{
		{
			SpaceComponent: engoc.SpaceComponent{
				Position: engo.Point{
					X: 200,
					Y: 200,
				},
				Width:  120,
				Height: 120,
			},
			animations: []*engoc.Animation{
				{
					Name:   "burn",
					Frames: continuous(61),
					Loop:   true,
				},
			},
			spritesheet: "bluefire.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
		{
			SpaceComponent: engoc.SpaceComponent{
				Position: engo.Point{
					X: 200,
					Y: 200,
				},
				Width:  120,
				Height: 120,
			},
			animations: []*engoc.Animation{
				{
					Name:   "burn",
					Frames: continuous(61),
					Loop:   true,
				},
			},
			spritesheet: "bluefire.png",
			width:       100,
			height:      100,
			rate:        0.1,
		},
	}
)

// Entity is a dynamic entity of world.
type Entity struct {
	ecs.BasicEntity

	engoc.AnimationComponent
	engoc.RenderComponent
	engoc.SpaceComponent

	// animations contains all entities animations informations for loading.
	// We use first element of slice to define default animation.
	animations []*engoc.Animation

	// Graphic informations
	spritesheet string  // filename of corresponding spreadsheet
	width       int     // width of one sprite
	height      int     // height of one sprite
	rate        float32 // Display rate of frames
}

func continuous(n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = i
	}
	return res
}

// NewEntity returns a new valid entity.
func NewEntity(e Entity) *Entity {
	e.BasicEntity = ecs.NewBasic()
	return &e
}

// LoadAnimations load all entity animations.
func (e *Entity) LoadAnimations() {

	spriteSheet := engoc.NewSpritesheetFromFile(e.spritesheet, e.width, e.height)

	e.RenderComponent.Drawable = spriteSheet.Cell(0)
	e.AnimationComponent = engoc.NewAnimationComponent(spriteSheet.Drawables(), e.rate)
	e.AnimationComponent.AddAnimations(e.animations)
	if len(e.animations) > 0 {
		e.AnimationComponent.AddDefaultAnimation(e.animations[0])
	}
}

func (e *Entity) AddToWorld(w *ecs.World) {
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *engoc.RenderSystem:
			sys.Add(&e.BasicEntity, &e.RenderComponent, &e.SpaceComponent)
		case *engoc.AnimationSystem:
			sys.Add(&e.BasicEntity, &e.AnimationComponent, &e.RenderComponent)
			// case *ControlSystem:
			// 	sys.Add(&ent.BasicEntity, &ent.AnimationComponent)
		}
	}
}
