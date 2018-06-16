package main

import (
	"sync"
	"testing"
	"time"

	nats "github.com/nats-io/go-nats"
	"github.com/stretchr/testify/assert"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/mocks"
	"github.com/elojah/game_01/storage"
)

func TestSequencer(t *testing.T) {

	equalEvent := func(lhs game.Event, rhs game.Event) bool {
		switch lhs.Action.(type) {
		case game.Cast:
			switch rhs.Action.(type) {
			case game.Cast:
				lhsTargets := lhs.Action.(game.Cast).Targets
				rhsTargets := rhs.Action.(game.Cast).Targets
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
	eset := []game.Event{
		game.Event{
			ID:     game.NewULID(),
			TS:     now,
			Action: game.Cast{Source: game.NewULID(), Targets: []game.ID{game.NewULID(), game.NewULID(), game.NewULID()}},
		},
		game.Event{
			ID:     game.NewULID(),
			TS:     now.Add(-1 * time.Second),
			Action: game.Move{Source: game.NewULID()},
		},
		game.Event{
			ID:     game.NewULID(),
			TS:     now.Add(-2 * time.Second),
			Action: game.Move{Source: game.NewULID()},
		},
		game.Event{
			ID:     game.NewULID(),
			TS:     now.Add(-3 * time.Second),
			Action: game.Move{Source: game.NewULID()},
		},
	}

	t.Run("simple", func(t *testing.T) {

		seqID := game.NewULID()
		es := mocks.NewEventMapper()
		es.ListEventFunc = func(subset game.EventSubset) ([]game.Event, error) {
			assert.Equal(t, seqID.String(), subset.Key)
			switch es.ListEventCount {
			case 0:
				assert.Equal(t, eset[0].TS.UnixNano(), subset.Min)
			}
			return []game.Event{eset[0]}, nil
		}

		var wg sync.WaitGroup
		wg.Add(1)
		seq := NewSequencer(seqID, 32, es,
			func(id game.ID, event game.Event) {
				assert.True(t, equalEvent(eset[0], event))
				wg.Done()
			},
		)
		seq.Start()

		raw, err := storage.NewEvent(eset[0]).Marshal(nil)
		assert.NoError(t, err)
		msg := &nats.Msg{Data: raw}
		seq.MsgHandler(msg)
		wg.Wait()
		seq.Close()
	})

	t.Run("two", func(t *testing.T) {

		seqID := game.NewULID()
		es := mocks.NewEventMapper()
		es.ListEventFunc = func(subset game.EventSubset) ([]game.Event, error) {
			assert.Equal(t, seqID.String(), subset.Key)
			switch int64(subset.Min) {
			case eset[0].TS.UnixNano():
				return []game.Event{eset[0]}, nil
			case eset[1].TS.UnixNano():
				return []game.Event{eset[1]}, nil
			}
			return nil, nil
		}

		var wg sync.WaitGroup
		wg.Add(2)
		seq := NewSequencer(seqID, 32, es,
			func(id game.ID, event game.Event) {
				wg.Done()
			},
		)
		seq.Start()

		raw1, err := storage.NewEvent(eset[1]).Marshal(nil)
		assert.NoError(t, err)
		msg1 := &nats.Msg{Data: raw1}
		seq.MsgHandler(msg1)

		raw0, err := storage.NewEvent(eset[0]).Marshal(nil)
		assert.NoError(t, err)
		msg0 := &nats.Msg{Data: raw0}
		seq.MsgHandler(msg0)

		wg.Wait()

		seq.Close()
	})

	t.Run("cancel", func(t *testing.T) {

		seqID := game.NewULID()
		es := mocks.NewEventMapper()
		es.ListEventFunc = func(subset game.EventSubset) ([]game.Event, error) {
			assert.Equal(t, seqID.String(), subset.Key)
			switch int64(subset.Min) {
			case eset[1].TS.UnixNano():
				return []game.Event{eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1],
					eset[0]}, nil
			case eset[2].TS.UnixNano():
				return []game.Event{eset[3], eset[2]}, nil
			}
			return nil, nil
		}

		var wg sync.WaitGroup
		wg.Add(1)
		seq := NewSequencer(seqID, 32, es,
			func(id game.ID, event game.Event) {
				assert.False(t, equalEvent(event, eset[0]))
				if equalEvent(event, eset[2]) {
					wg.Done()
				}
			},
		)
		seq.Start()

		raw1, err := storage.NewEvent(eset[1]).Marshal(nil)
		assert.NoError(t, err)
		msg1 := &nats.Msg{Data: raw1}
		seq.MsgHandler(msg1)

		raw2, err := storage.NewEvent(eset[2]).Marshal(nil)
		assert.NoError(t, err)
		msg2 := &nats.Msg{Data: raw2}
		seq.MsgHandler(msg2)

		wg.Wait()

		seq.Close()
	})

	t.Run("interrupt", func(t *testing.T) {

		seqID := game.NewULID()
		es := mocks.NewEventMapper()
		es.ListEventFunc = func(subset game.EventSubset) ([]game.Event, error) {
			assert.Equal(t, seqID.String(), subset.Key)
			switch int64(subset.Min) {
			case eset[1].TS.UnixNano():
				time.Sleep(10 * time.Millisecond)
				return []game.Event{eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1], eset[1],
					eset[0]}, nil
			case eset[2].TS.UnixNano():
				return []game.Event{eset[3], eset[2]}, nil
			}
			return nil, nil
		}

		var wg sync.WaitGroup
		wg.Add(1)
		seq := NewSequencer(seqID, 1, es,
			func(id game.ID, event game.Event) {
				assert.False(t, equalEvent(event, eset[0]))
				if equalEvent(event, eset[2]) {
					wg.Done()
				}
			},
		)
		seq.Start()

		raw1, err := storage.NewEvent(eset[1]).Marshal(nil)
		assert.NoError(t, err)
		msg1 := &nats.Msg{Data: raw1}
		seq.MsgHandler(msg1)

		time.Sleep(10*time.Millisecond + 1*time.Nanosecond)

		raw2, err := storage.NewEvent(eset[2]).Marshal(nil)
		assert.NoError(t, err)
		msg2 := &nats.Msg{Data: raw2}
		seq.MsgHandler(msg2)

		wg.Wait()

		seq.Close()
	})

}
