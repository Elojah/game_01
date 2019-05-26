package app

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elojah/game_01/pkg/ability"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	eventmocks "github.com/elojah/game_01/pkg/event/mocks"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

func TestA(t *testing.T) {

	t.Run("set event perform source success", func(t *testing.T) {

		// Data
		e := event.E{
			ID:    gulid.NewID(),
			Token: gulid.NewID(),
			Action: event.Action{
				PerformSource: &event.PerformSource{
					Targets: map[string]ability.Targets{
						"test": {
							Entities: gulid.NewIDs(3),
						},
					},
				},
			},
			Trigger: gulid.NewID(),
		}
		entityID := gulid.NewID()

		// Mocks
		store := &eventmocks.Store{
			UpsertFunc: func(newE event.E, id gulid.ID) error {
				assert.Equal(t, e.ID.String(), newE.ID.String())
				assert.Equal(t, entityID.String(), id.String())
				return nil
			},
		}
		triggerStore := &eventmocks.TriggerStore{
			FetchTriggerFunc: func(triggerID gulid.ID, id gulid.ID) (gulid.ID, error) {
				assert.Equal(t, e.Trigger.String(), triggerID.String())
				assert.Equal(t, entityID.String(), id.String())
				return gulid.ID{}, gerrors.ErrNotFound{Store: "test_trigger", Index: triggerID.String()}
			},
			UpsertTriggerFunc: func(trigger event.Trigger) error {
				assert.Equal(t, entityID.String(), trigger.EntityID.String())
				assert.Equal(t, e.Trigger.String(), trigger.EventSourceID.String())
				assert.Equal(t, e.ID.String(), trigger.EventTargetID.String())
				return nil
			},
		}
		qStore := &eventmocks.QStore{}

		// Test
		app := A{
			TriggerStore: triggerStore,
			Store:        store,
			QStore:       qStore,
		}
		assert.NoError(t, app.Create(e, entityID))

		// Assert
		assert.Equal(t, int32(1), store.UpsertCount)
		assert.Equal(t, int32(0), store.FetchCount)
		assert.Equal(t, int32(0), store.ListCount)
		assert.Equal(t, int32(0), store.RemoveCount)

		assert.Equal(t, int32(1), triggerStore.UpsertTriggerCount)
		assert.Equal(t, int32(1), triggerStore.FetchTriggerCount)
		assert.Equal(t, int32(0), triggerStore.ListTriggerCount)
		assert.Equal(t, int32(0), triggerStore.RemoveTriggerCount)

		assert.Equal(t, int32(0), qStore.PublishCount)
		assert.Equal(t, int32(0), qStore.SubscribeCount)
	})

	t.Run("cancel event success", func(t *testing.T) {

		// Data
		e := event.E{
			ID:    gulid.NewID(),
			Token: gulid.NewID(),
			Action: event.Action{
				Cancel: &event.Cancel{},
			},
			Trigger: gulid.NewID(),
		}
		prevE := event.E{
			ID:    gulid.NewID(),
			Token: e.Token,
			Action: event.Action{
				PerformSource: &event.PerformSource{
					Targets: map[string]ability.Targets{
						"test": {
							Entities: gulid.NewIDs(2),
						},
					},
				},
			},
			Trigger: e.Trigger,
		}
		entityID := gulid.NewID()

		// Mocks
		store := &eventmocks.Store{
			FetchFunc: func(prevID gulid.ID, id gulid.ID) (event.E, error) {
				assert.Equal(t, prevE.ID.String(), prevID.String())
				assert.Equal(t, entityID.String(), id.String())
				return prevE, nil
			},
			RemoveFunc: func(prevID gulid.ID, id gulid.ID) error {
				assert.Equal(t, prevE.ID.String(), prevID.String())
				assert.Equal(t, entityID.String(), id.String())
				return nil
			},
		}
		triggerStore := &eventmocks.TriggerStore{
			FetchTriggerFunc: func(triggerID gulid.ID, id gulid.ID) (gulid.ID, error) {
				assert.Equal(t, e.Trigger.String(), triggerID.String())
				assert.Equal(t, entityID.String(), id.String())
				return prevE.ID, nil
			},
			RemoveTriggerFunc: func(triggerID gulid.ID, id gulid.ID) error {
				assert.Equal(t, e.Trigger.String(), triggerID.String())
				assert.Equal(t, entityID.String(), id.String())
				return nil
			},
		}
		qStore := &eventmocks.QStore{
			PublishFunc: func(cancelE event.E, id gulid.ID) error {
				assert.NotNil(t, e.Action.Cancel)
				return nil
			},
		}

		// Test
		app := A{
			TriggerStore: triggerStore,
			Store:        store,
			QStore:       qStore,
		}
		assert.NoError(t, app.Create(e, entityID))

		// Assert
		assert.Equal(t, int32(0), store.UpsertCount)
		assert.Equal(t, int32(1), store.FetchCount)
		assert.Equal(t, int32(0), store.ListCount)
		assert.Equal(t, int32(1), store.RemoveCount)

		assert.Equal(t, int32(0), triggerStore.UpsertTriggerCount)
		assert.Equal(t, int32(1), triggerStore.FetchTriggerCount)
		assert.Equal(t, int32(0), triggerStore.ListTriggerCount)
		assert.Equal(t, int32(1), triggerStore.RemoveTriggerCount)

		assert.Equal(t, int32(2), qStore.PublishCount)
		assert.Equal(t, int32(0), qStore.SubscribeCount)
	})

}
