package ui

import (
	"github.com/EngoEngine/ecs"
	engoc "github.com/EngoEngine/engo/common"
	"github.com/elojah/game_01/pkg/ulid"
)

// Entity is a dynamic entity of world.
type Entity struct {
	GID ulid.ID

	Player bool // player indicates if this entity is player controlled

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

func continuous(start int, end int) []int {
	if end < start {
		return nil
	}
	res := make([]int, end-start)
	for i := 0; i < end-start; i++ {
		res[i] = i + start
	}
	return res
}

// NewEntity returns a new valid entity.
func NewEntity(e Entity) *Entity {
	e.GID = ulid.NewID()
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
		case *ControlSystem:
			if e.Player {
				sys.Add(e)
			}
		}
	}
}

func (e *Entity) RemoveFromWorld(w *ecs.World) {
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *engoc.RenderSystem:
			sys.Remove(e.BasicEntity)
		case *engoc.AnimationSystem:
			sys.Remove(e.BasicEntity)
		case *ControlSystem:
			if e.Player {
				sys.Remove(e.BasicEntity)
			}
		}
	}
}
