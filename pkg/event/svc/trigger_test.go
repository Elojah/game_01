package svc

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elojah/game_01/pkg/ability"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	eventmocks "github.com/elojah/game_01/pkg/event/mocks"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

func TestTriggerService(t *testing.T) {

	t.Run("set event move source success", func(t *testing.T) {

		// Data
		e := event.E{
			ID:    gulid.NewID(),
			Token: gulid.NewID(),
			Action: event.Action{
				CastSource: &event.CastSource{
					Targets: map[string]ability.Targets{
						"test": ability.Targets{
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
			SetEventFunc: func(newE event.E, id gulid.ID) error {
				assert.Equal(t, e.ID.String(), newE.ID.String())
				assert.Equal(t, entityID.String(), id.String())
				return nil
			},
		}
		triggerStore := &eventmocks.TriggerStore{
			GetTriggerFunc: func(triggerID gulid.ID, id gulid.ID) (gulid.ID, error) {
				assert.Equal(t, e.Trigger.String(), triggerID.String())
				assert.Equal(t, entityID.String(), id.String())
				return gulid.ID{}, gerrors.ErrNotFound{Store: "test_trigger", Index: triggerID.String()}
			},
			SetTriggerFunc: func(trigger event.Trigger) error {
				assert.Equal(t, entityID.String(), trigger.EntityID.String())
				assert.Equal(t, e.Trigger.String(), trigger.EventSourceID.String())
				assert.Equal(t, e.ID.String(), trigger.EventTargetID.String())
				return nil
			},
		}
		qStore := &eventmocks.QStore{}

		// Test
		s := TriggerService{
			TriggerStore: triggerStore,
			Store:        store,
			QStore:       qStore,
		}
		assert.NoError(t, s.Set(e, entityID))

		// Assert
		assert.Equal(t, int32(1), store.SetEventCount)
		assert.Equal(t, int32(0), store.GetEventCount)
		assert.Equal(t, int32(0), store.ListEventCount)
		assert.Equal(t, int32(0), store.DelEventCount)

		assert.Equal(t, int32(1), triggerStore.SetTriggerCount)
		assert.Equal(t, int32(1), triggerStore.GetTriggerCount)
		assert.Equal(t, int32(0), triggerStore.ListTriggerCount)
		assert.Equal(t, int32(0), triggerStore.DelTriggerCount)

		assert.Equal(t, int32(0), qStore.PublishEventCount)
		assert.Equal(t, int32(0), qStore.SubscribeEventCount)
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
				CastSource: &event.CastSource{
					Targets: map[string]ability.Targets{
						"test": ability.Targets{
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
			GetEventFunc: func(prevID gulid.ID, id gulid.ID) (event.E, error) {
				assert.Equal(t, prevE.ID.String(), prevID.String())
				assert.Equal(t, entityID.String(), id.String())
				return prevE, nil
			},
			DelEventFunc: func(prevID gulid.ID, id gulid.ID) error {
				assert.Equal(t, prevE.ID.String(), prevID.String())
				assert.Equal(t, entityID.String(), id.String())
				return nil
			},
		}
		triggerStore := &eventmocks.TriggerStore{
			GetTriggerFunc: func(triggerID gulid.ID, id gulid.ID) (gulid.ID, error) {
				assert.Equal(t, e.Trigger.String(), triggerID.String())
				assert.Equal(t, entityID.String(), id.String())
				return prevE.ID, nil
			},
			DelTriggerFunc: func(triggerID gulid.ID, id gulid.ID) error {
				assert.Equal(t, e.Trigger.String(), triggerID.String())
				assert.Equal(t, entityID.String(), id.String())
				return nil
			},
		}
		qStore := &eventmocks.QStore{
			PublishEventFunc: func(cancelE event.E, id gulid.ID) error {
				assert.NotNil(t, e.Action.Cancel)
				return nil
			},
		}

		// Test
		s := TriggerService{
			TriggerStore: triggerStore,
			Store:        store,
			QStore:       qStore,
		}
		assert.NoError(t, s.Set(e, entityID))

		// Assert
		assert.Equal(t, int32(0), store.SetEventCount)
		assert.Equal(t, int32(1), store.GetEventCount)
		assert.Equal(t, int32(0), store.ListEventCount)
		assert.Equal(t, int32(1), store.DelEventCount)

		assert.Equal(t, int32(0), triggerStore.SetTriggerCount)
		assert.Equal(t, int32(1), triggerStore.GetTriggerCount)
		assert.Equal(t, int32(0), triggerStore.ListTriggerCount)
		assert.Equal(t, int32(1), triggerStore.DelTriggerCount)

		assert.Equal(t, int32(2), qStore.PublishEventCount)
		assert.Equal(t, int32(0), qStore.SubscribeEventCount)
	})

}
