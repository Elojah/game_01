package main

import (
	"sync"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	entitymocks "github.com/elojah/game_01/pkg/entity/mocks"
	"github.com/elojah/game_01/pkg/event"
	eventmocks "github.com/elojah/game_01/pkg/event/mocks"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
)

func TestSequencer(t *testing.T) {

	now := time.Now()
	eset := []event.E{
		event.E{
			ID: ulid.NewID(),
			TS: now,
			Action: event.Action{
				Cast: &event.Cast{Source: ulid.NewID(), Targets: []ulid.ID{ulid.NewID(), ulid.NewID(), ulid.NewID()}},
			},
		},
		event.E{
			ID: ulid.NewID(),
			TS: now.Add(-1 * time.Second),
			Action: event.Action{
				Move: &event.Move{Source: ulid.NewID()},
			},
		},
		event.E{
			ID: ulid.NewID(),
			TS: now.Add(-2 * time.Second),
			Action: event.Action{
				Move: &event.Move{Source: ulid.NewID()},
			},
		},
		event.E{
			ID: ulid.NewID(),
			TS: now.Add(-3 * time.Second),
			Action: event.Action{
				Move: &event.Move{Source: ulid.NewID()},
			},
		},
	}

	t.Run("simple", func(t *testing.T) {

		seqID := ulid.NewID()
		entityStore := entitymocks.NewStore()
		eventStore := eventmocks.NewStore()
		eventStore.ListEventFunc = func(subset event.Subset) ([]event.E, error) {
			assert.Equal(t, seqID.String(), subset.Key)
			switch eventStore.ListEventCount {
			case 0:
				assert.Equal(t, eset[0].TS.UnixNano(), subset.Min)
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
			switch int64(subset.Min) {
			case eset[0].TS.UnixNano():
				return []event.E{eset[0]}, nil
			case eset[1].TS.UnixNano():
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

		raw1, err := eset[1].Marshal()
		assert.NoError(t, err)
		msg1 := &infra.Message{Payload: string(raw1)}
		seq.Handler(msg1)

		raw0, err := eset[0].Marshal()
		assert.NoError(t, err)
		msg0 := &infra.Message{Payload: string(raw0)}
		seq.Handler(msg0)

		wg.Wait()

		seq.Close()
	})

	t.Run("cancel", func(t *testing.T) {

		seqID := ulid.NewID()
		entityStore := entitymocks.NewStore()
		eventStore := eventmocks.NewStore()
		eventStore.ListEventFunc = func(subset event.Subset) ([]event.E, error) {
			assert.Equal(t, seqID.String(), subset.Key)
			switch int64(subset.Min) {
			case eset[1].TS.UnixNano():
				return []event.E{eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1],
					eset[0]}, nil
			case eset[2].TS.UnixNano():
				return []event.E{eset[3], eset[2]}, nil
			}
			return nil, nil
		}

		var wg sync.WaitGroup
		wg.Add(1)
		seq := NewSequencer(seqID, 32,
			func(id ulid.ID, e event.E) {
				assert.False(t, e.Equal(eset[0]))
				if e.Equal(eset[2]) {
					wg.Done()
				}
			},
		)
		seq.EventStore = eventStore
		seq.EntityStore = entityStore
		seq.logger = zerolog.Nop()
		seq.Run()

		raw1, err := eset[1].Marshal()
		assert.NoError(t, err)
		msg1 := &infra.Message{Payload: string(raw1)}
		seq.Handler(msg1)

		raw2, err := eset[2].Marshal()
		assert.NoError(t, err)
		msg2 := &infra.Message{Payload: string(raw2)}
		seq.Handler(msg2)

		wg.Wait()

		seq.Close()
	})

	t.Run("interrupt", func(t *testing.T) {

		seqID := ulid.NewID()
		entityStore := entitymocks.NewStore()
		eventStore := eventmocks.NewStore()
		eventStore.ListEventFunc = func(subset event.Subset) ([]event.E, error) {
			assert.Equal(t, seqID.String(), subset.Key)
			switch int64(subset.Min) {
			case eset[1].TS.UnixNano():
				time.Sleep(10 * time.Millisecond)
				return []event.E{eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1],
					eset[0]}, nil
			case eset[2].TS.UnixNano():
				return []event.E{eset[3], eset[2]}, nil
			}
			return nil, nil
		}

		var wg sync.WaitGroup
		wg.Add(1)
		seq := NewSequencer(seqID, 1,
			func(id ulid.ID, e event.E) {
				assert.False(t, e.Equal(eset[0]))
				if e.Equal(eset[2]) {
					wg.Done()
				}
			},
		)
		seq.EventStore = eventStore
		seq.EntityStore = entityStore
		seq.logger = zerolog.Nop()
		seq.Run()

		raw1, err := eset[1].Marshal()
		assert.NoError(t, err)
		msg1 := &infra.Message{Payload: string(raw1)}
		seq.Handler(msg1)

		time.Sleep(10*time.Millisecond + 1*time.Nanosecond)

		raw2, err := eset[2].Marshal()
		assert.NoError(t, err)
		msg2 := &infra.Message{Payload: string(raw2)}
		seq.Handler(msg2)

		wg.Wait()

		seq.Close()
	})

}
