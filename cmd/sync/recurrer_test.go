package main

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"

	game "github.com/elojah/game_01"
	"github.com/elojah/game_01/mocks"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/sector"
)

func TestRecurrer(t *testing.T) {

	entities := []entity.E{
		entity.E{
			ID:   game.NewID(),
			Type: game.NewID(),
			Name: "La muerte del sol",
			HP:   76,
			MP:   567,
			Position: entity.Position{
				SectorID: game.NewID(),
				Coord:    game.Vec3{X: 53.1233, Y: 68.0706, Z: 67.0753},
			},
		},
	}

	t.Run("simple", func(t *testing.T) {
		var count int32
		var wg sync.WaitGroup
		entityMock := &mocks.EntityMapper{
			GetEntityFunc: func(subset entity.Subset) (entity.E, error) {
				return entities[0], nil
			},
		}
		sectorMock := &mocks.SectorMapper{
			GetSectorFunc: func(subset sector.Subset) (sector.S, error) {
				return sector.S{ID: entities[0].Position.SectorID}, nil
			},
		}
		sectorEntitiesMock := &mocks.SectorEntitiesMapper{
			GetEntitiesFunc: func(subset sector.EntitiesSubset) (sector.Entities, error) {
				return sector.Entities{
					SectorID:  entities[0].Position.SectorID,
					EntityIDs: []game.ID{entities[0].ID},
				}, nil
			},
		}
		wg.Add(1)
		callback := func(e entity.E) {
			atomic.AddInt32(&count, 1)
			assert.True(t, entities[0].Equal(e))
			wg.Done()
		}
		r := event.Recurrer{ID: game.NewID(), EntityID: game.NewID(), TokenID: game.NewID()}
		rec := NewRecurrer(r, 10, callback)
		rec.EntityMapper = entityMock
		rec.SectorMapper = sectorMock
		rec.EntitiesMapper = sectorEntitiesMock
		go rec.Start()
		defer rec.Close()
		wg.Wait()
	})
}
