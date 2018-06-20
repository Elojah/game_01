package main

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"

	game "github.com/elojah/game_01"
	"github.com/elojah/game_01/mocks"
)

func TestRecurrer(t *testing.T) {

	entities := []game.Entity{
		game.Entity{
			ID:   game.NewID(),
			Type: game.NewID(),
			Name: "La muerte del sol",
			HP:   76,
			MP:   567,
			Position: game.Position{
				SectorID: game.NewID(),
				Coord:    game.Vec3{X: 53.1233, Y: 68.0706, Z: 67.0753},
			},
		},
	}

	t.Run("simple", func(t *testing.T) {
		var count int32
		var wg sync.WaitGroup
		entityMock := &mocks.EntityMapper{
			GetEntityFunc: func(subset game.EntitySubset) (game.Entity, error) {
				return entities[0], nil
			},
		}
		sectorMock := &mocks.SectorMapper{
			GetSectorFunc: func(subset game.SectorSubset) (game.Sector, error) {
				return game.Sector{ID: entities[0].Position.SectorID}, nil
			},
		}
		sectorEntitiesMock := &mocks.SectorEntitiesMapper{
			GetSectorEntitiesFunc: func(subset game.SectorEntitiesSubset) (game.SectorEntities, error) {
				return game.SectorEntities{
					SectorID:  entities[0].Position.SectorID,
					EntityIDs: []game.ID{entities[0].ID},
				}, nil
			},
		}
		wg.Add(1)
		callback := func(entity game.Entity) {
			atomic.AddInt32(&count, 1)
			assert.True(t, entities[0].Equal(entity))
			wg.Done()
		}
		r := game.Recurrer{ID: game.NewID(), EntityID: game.NewID(), TokenID: game.NewID()}
		rec := NewRecurrer(r, 10, callback)
		rec.EntityMapper = entityMock
		rec.SectorMapper = sectorMock
		rec.SectorEntitiesMapper = sectorEntitiesMock
		go rec.Start()
		defer rec.Close()
		wg.Wait()
	})
}
