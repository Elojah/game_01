package main

import (
	"sync"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/mocks"
	"github.com/elojah/game_01/pkg/storage"
	"github.com/elojah/game_01/pkg/ulid"
)

func TestSequencer(t *testing.T) {

	equalEvent := func(lhs event.E, rhs event.E) bool {
		switch lhs.Action.(type) {
		case event.Cast:
			switch rhs.Action.(type) {
			case event.Cast:
				lhsTargets := lhs.Action.(event.Cast).Targets
				rhsTargets := rhs.Action.(event.Cast).Targets
				for i, target := range lhsTargets {
					if target.Compare(rhsTargets[i]) != 0 {
						return false
					}
				}
			default:
				return false
			}
		default:
			if lhs.Action != rhs.Action {
				return false
			}
		}
		return lhs.ID.Compare(rhs.ID) == 0 &&
			lhs.TS.Equal(rhs.TS)
	}

	now := time.Now()
	eset := []event.E{
		event.E{
			ID:     ulid.NewID(),
			TS:     now,
			Action: event.Cast{Source: ulid.NewID(), Targets: []ulid.ID{ulid.NewID(), ulid.NewID(), ulid.NewID()}},
		},
		event.E{
			ID:     ulid.NewID(),
			TS:     now.Add(-1 * time.Second),
			Action: event.Move{Source: ulid.NewID()},
		},
		event.E{
			ID:     ulid.NewID(),
			TS:     now.Add(-2 * time.Second),
			Action: event.Move{Source: ulid.NewID()},
		},
		event.E{
			ID:     ulid.NewID(),
			TS:     now.Add(-3 * time.Second),
			Action: event.Move{Source: ulid.NewID()},
		},
	}

	t.Run("simple", func(t *testing.T) {

		seqID := ulid.NewID()
		es := mocks.NewEventMapper()
		es.ListEventFunc = func(subset event.Subset) ([]event.E, error) {
			assert.Equal(t, seqID.String(), subset.Key)
			switch es.ListEventCount {
			case 0:
				assert.Equal(t, eset[0].TS.UnixNano(), subset.Min)
			}
			return []event.E{eset[0]}, nil
		}

		var wg sync.WaitGroup
		wg.Add(1)
		seq := NewSequencer(seqID, 32, es,
			func(id ulid.ID, e event.E) {
				assert.True(t, equalEvent(eset[0], e))
				wg.Done()
			},
		)
		seq.logger = zerolog.Nop()
		seq.Run()

		raw, err := storage.NewEvent(eset[0]).Marshal(nil)
		assert.NoError(t, err)
		msg := &event.Message{Payload: string(raw)}
		seq.Handler(msg)
		wg.Wait()
		seq.Close()
	})

	t.Run("two", func(t *testing.T) {

		seqID := ulid.NewID()
		es := mocks.NewEventMapper()
		es.ListEventFunc = func(subset event.Subset) ([]event.E, error) {
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
		seq := NewSequencer(seqID, 32, es,
			func(id ulid.ID, e event.E) {
				wg.Done()
			},
		)
		seq.logger = zerolog.Nop()
		seq.Run()

		raw1, err := storage.NewEvent(eset[1]).Marshal(nil)
		assert.NoError(t, err)
		msg1 := &event.Message{Payload: string(raw1)}
		seq.Handler(msg1)

		raw0, err := storage.NewEvent(eset[0]).Marshal(nil)
		assert.NoError(t, err)
		msg0 := &event.Message{Payload: string(raw0)}
		seq.Handler(msg0)

		wg.Wait()

		seq.Close()
	})

	t.Run("cancel", func(t *testing.T) {

		seqID := ulid.NewID()
		es := mocks.NewEventMapper()
		es.ListEventFunc = func(subset event.Subset) ([]event.E, error) {
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
		seq := NewSequencer(seqID, 32, es,
			func(id ulid.ID, e event.E) {
				assert.False(t, equalEvent(e, eset[0]))
				if equalEvent(e, eset[2]) {
					wg.Done()
				}
			},
		)
		seq.logger = zerolog.Nop()
		seq.Run()

		raw1, err := storage.NewEvent(eset[1]).Marshal(nil)
		assert.NoError(t, err)
		msg1 := &event.Message{Payload: string(raw1)}
		seq.Handler(msg1)

		raw2, err := storage.NewEvent(eset[2]).Marshal(nil)
		assert.NoError(t, err)
		msg2 := &event.Message{Payload: string(raw2)}
		seq.Handler(msg2)

		wg.Wait()

		seq.Close()
	})

	t.Run("interrupt", func(t *testing.T) {

		seqID := ulid.NewID()
		es := mocks.NewEventMapper()
		es.ListEventFunc = func(subset event.Subset) ([]event.E, error) {
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
		seq := NewSequencer(seqID, 1, es,
			func(id ulid.ID, e event.E) {
				assert.False(t, equalEvent(e, eset[0]))
				if equalEvent(e, eset[2]) {
					wg.Done()
				}
			},
		)
		seq.logger = zerolog.Nop()
		seq.Run()

		raw1, err := storage.NewEvent(eset[1]).Marshal(nil)
		assert.NoError(t, err)
		msg1 := &event.Message{Payload: string(raw1)}
		seq.Handler(msg1)

		time.Sleep(10*time.Millisecond + 1*time.Nanosecond)

		raw2, err := storage.NewEvent(eset[2]).Marshal(nil)
		assert.NoError(t, err)
		msg2 := &event.Message{Payload: string(raw2)}
		seq.Handler(msg2)

		wg.Wait()

		seq.Close()
	})

}
