package main

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/elojah/game_01/pkg/ability"
	entitymocks "github.com/elojah/game_01/pkg/entity/mocks"
	"github.com/elojah/game_01/pkg/event"
	eventmocks "github.com/elojah/game_01/pkg/event/mocks"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
)

func TestSequencer(t *testing.T) {

	now := time.Now()
	cid := ulid.NewID().String()
	eset := []event.E{
		event.E{
			ID: ulid.NewTimeID(now.Unix()),
			Action: event.Action{
				CastSource: &event.CastSource{
					AbilityID: ulid.NewID(),
					Targets: map[string]ability.Targets{
						cid: ability.Targets{
							Entities: []ulid.ID{ulid.NewID(), ulid.NewID(), ulid.NewID()},
						},
					},
				},
			},
		},
		event.E{
			ID: ulid.NewTimeID(now.Add(1 * time.Second).Unix()),
			Action: event.Action{
				MoveTarget: &event.MoveTarget{Source: ulid.NewID()},
			},
		},
		event.E{
			ID: ulid.NewTimeID(now.Add(2 * time.Second).Unix()),
			Action: event.Action{
				MoveTarget: &event.MoveTarget{Source: ulid.NewID()},
			},
		},
		event.E{
			ID: ulid.NewTimeID(now.Add(3 * time.Second).Unix()),
			Action: event.Action{
				MoveTarget: &event.MoveTarget{Source: ulid.NewID()},
			},
		},
	}
	fmt.Println(0, eset[0].ID.String())
	fmt.Println(1, eset[1].ID.String())
	fmt.Println(2, eset[2].ID.String())
	fmt.Println(3, eset[3].ID.String())

	t.Run("simple", func(t *testing.T) {

		seqID := ulid.NewID()
		entityStore := entitymocks.NewStore()
		eventStore := eventmocks.NewStore()
		eventStore.ListEventFunc = func(subset event.Subset) ([]event.E, error) {
			assert.Equal(t, seqID.String(), subset.Key)
			switch eventStore.ListEventCount {
			case 0:
				assert.Equal(t, eset[0].ID.String(), subset.Min.String())
			}
			return []event.E{eset[0]}, nil
		}

		var wg sync.WaitGroup
		wg.Add(1)
		seq := NewSequencer(seqID, 32,
			func(id ulid.ID, e event.E) {
				assert.True(t, eset[0].Equal(e))
				wg.Done()
			},
		)
		seq.EventStore = eventStore
		seq.EntityStore = entityStore
		seq.logger = zerolog.Nop()
		seq.Run()

		raw, err := eset[0].Marshal()
		assert.NoError(t, err)
		msg := &infra.Message{Payload: string(raw)}
		seq.Handler(msg)
		wg.Wait()
		seq.Close()
	})

	t.Run("two", func(t *testing.T) {

		seqID := ulid.NewID()
		entityStore := entitymocks.NewStore()
		eventStore := eventmocks.NewStore()
		eventStore.ListEventFunc = func(subset event.Subset) ([]event.E, error) {
			assert.Equal(t, seqID.String(), subset.Key)
			switch subset.Min.String() {
			case eset[0].ID.String():
				return []event.E{eset[0]}, nil
			case eset[1].ID.String():
				return []event.E{eset[1]}, nil
			}
			return nil, nil
		}

		var wg sync.WaitGroup
		wg.Add(2)
		seq := NewSequencer(seqID, 32,
			func(id ulid.ID, e event.E) {
				wg.Done()
			},
		)
		seq.EventStore = eventStore
		seq.EntityStore = entityStore
		seq.logger = zerolog.Nop()
		seq.Run()

		raw0, err := eset[0].Marshal()
		assert.NoError(t, err)
		msg0 := &infra.Message{Payload: string(raw0)}

		raw1, err := eset[1].Marshal()
		assert.NoError(t, err)
		msg1 := &infra.Message{Payload: string(raw1)}

		seq.Handler(msg0)
		seq.Handler(msg1)

		wg.Wait()

		seq.Close()
	})

	t.Run("cancel", func(t *testing.T) {

		seqID := ulid.NewID()
		entityStore := entitymocks.NewStore()
		eventStore := eventmocks.NewStore()
		eventStore.ListEventFunc = func(subset event.Subset) ([]event.E, error) {
			assert.Equal(t, seqID.String(), subset.Key)
			switch subset.Min.String() {
			case eset[0].ID.String():
				return []event.E{eset[0], eset[1]}, nil
			case eset[2].ID.String():
				n := 42
				es := make([]event.E, n)
				for i := 0; i < n; i++ {
					es[i] = eset[2]
				}
				return append(es, eset[3]), nil
			}
			return nil, nil
		}

		var wg sync.WaitGroup
		wg.Add(1)
		seq := NewSequencer(seqID, 32,
			func(id ulid.ID, e event.E) {
				fmt.Println(e.ID.String())
				assert.False(t, e.Equal(eset[3]))
				if e.Equal(eset[1]) {
					wg.Done()
				}
			},
		)
		seq.EventStore = eventStore
		seq.EntityStore = entityStore
		seq.logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
		// seq.logger = zerolog.Nop()
		seq.Run()

		raw2, err := eset[2].Marshal()
		assert.NoError(t, err)
		msg2 := &infra.Message{Payload: string(raw2)}

		raw0, err := eset[0].Marshal()
		assert.NoError(t, err)
		msg0 := &infra.Message{Payload: string(raw0)}

		seq.Handler(msg2)
		seq.Handler(msg0)

		wg.Wait()

		seq.Close()
	})

	t.Run("interrupt", func(t *testing.T) {

		seqID := ulid.NewID()
		entityStore := entitymocks.NewStore()
		eventStore := eventmocks.NewStore()
		eventStore.ListEventFunc = func(subset event.Subset) ([]event.E, error) {
			assert.Equal(t, seqID.String(), subset.Key)
			switch subset.Min.String() {
			case eset[0].ID.String():
				return []event.E{eset[0], eset[1]}, nil
			case eset[2].ID.String():
				n := 4242
				es := make([]event.E, n)
				for i := 0; i < n; i++ {
					es[i] = eset[2]
				}
				return append(es, eset[3]), nil
			}
			return nil, nil
		}

		var wg sync.WaitGroup
		wg.Add(1)
		seq := NewSequencer(seqID, 1,
			func(id ulid.ID, e event.E) {
				if e.Equal(eset[2]) {
					time.Sleep(5 * time.Millisecond)
				}
				assert.False(t, e.Equal(eset[3]))
				if e.Equal(eset[1]) {
					wg.Done()
				}
			},
		)
		seq.EventStore = eventStore
		seq.EntityStore = entityStore
		seq.logger = zerolog.Nop()
		seq.Run()

		raw2, err := eset[2].Marshal()
		assert.NoError(t, err)
		msg2 := &infra.Message{Payload: string(raw2)}

		raw0, err := eset[0].Marshal()
		assert.NoError(t, err)
		msg0 := &infra.Message{Payload: string(raw0)}

		seq.Handler(msg2)
		time.Sleep(10 * time.Millisecond)
		seq.Handler(msg0)

		wg.Wait()

		seq.Close()
	})

}
