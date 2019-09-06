package ui

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	engoc "github.com/EngoEngine/engo/common"
)

// Tile  represents a square tile map.
type Tile struct {
	ecs.BasicEntity
	engoc.RenderComponent
	engoc.SpaceComponent
}

// Tiles alias a slice of Tile.
type Tiles []Tile

// New loads tiles from a map tmx file.
func NewTiles(filename string) (Tiles, error) {

	resource, err := engo.Files.Resource(filename)
	if err != nil {
		return nil, err
	}

	tmxResource := resource.(engoc.TMXResource)
	level := tmxResource.Level

	var ts Tiles
	for _, tlayer := range level.TileLayers {
		for _, elem := range tlayer.Tiles {
			if elem.Image != nil {
				tile := Tile{BasicEntity: ecs.NewBasic()}
				tile.RenderComponent = engoc.RenderComponent{
					Drawable: elem,
					Scale:    engo.Point{1, 1},
				}
				tile.SpaceComponent = engoc.SpaceComponent{
					Position: elem.Point,
					Width:    0,
					Height:   0,
				}
				ts = append(ts, tile)
			}
		}
	}
	return ts, nil
}

// AddToWorld adds tiles map to world.
func (ts Tiles) AddToWorld(w *ecs.World) {
	// add the ts to the RenderSystem
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *engoc.RenderSystem:
			for _, v := range ts {
				sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
			}
		}
	}
}
