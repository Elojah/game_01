package main

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elojah/game_01/mocks"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/geometry"
	"github.com/elojah/game_01/pkg/sector"
	"github.com/elojah/game_01/pkg/ulid"
)

func TestRecurrer(t *testing.T) {

	entities := []entity.E{
		entity.E{
			ID:   ulid.NewID(),
			Type: ulid.NewID(),
			Name: "La muerte del sol",
			HP:   76,
			MP:   567,
			Position: entity.Position{
				SectorID: ulid.NewID(),
				Coord:    geometry.Vec3{X: 53.1233, Y: 68.0706, Z: 67.0753},
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
					EntityIDs: []ulid.ID{entities[0].ID},
				}, nil
			},
		}
		wg.Add(1)
		callback := func(e entity.E) {
			atomic.AddInt32(&count, 1)
			assert.True(t, entities[0].Equal(e))
			wg.Done()
		}
		r := event.Recurrer{ID: ulid.NewID(), EntityID: ulid.NewID(), TokenID: ulid.NewID()}
		rec := NewRecurrer(r, 10, callback)
		rec.EntityMapper = entityMock
		rec.SectorMapper = sectorMock
		rec.EntitiesMapper = sectorEntitiesMock
		go rec.Start()
		defer rec.Close()
		wg.Wait()
	})
}
