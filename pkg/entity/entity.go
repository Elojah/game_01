package entity

import (
	multierror "github.com/hashicorp/go-multierror"

	"github.com/elojah/game_01/pkg/ability"
	gerrors "github.com/elojah/game_01/pkg/errors"
	gulid "github.com/elojah/game_01/pkg/ulid"
)

// Store contains basic operations for entity E object.
type Store interface {
	Upsert(E, uint64) error
	Fetch(gulid.ID, uint64) (E, error)
	Remove(gulid.ID) error
	RemoveByTS(gulid.ID, uint64) error
}

// App contains entity stores and applications.
type App interface {
	InventoryStore
	MRInventoryStore
	PCLeftStore
	PCStore
	PermissionStore
	SpawnStore
	Store
	TemplateStore

	Disconnect(id gulid.ID) error
	ErasePC(gulid.ID, gulid.ID) error
	FetchMRInventoryFromCache(gulid.ID, gulid.ID) (Inventory, error)
	UpsertMRInventoryWithCache(gulid.ID, Inventory) error
	CheckPermission(gulid.ID, gulid.ID) error
}

// CastAbility decreases MP (without check) and assign a new cast to entity.
func (e *E) CastAbility(ab ability.A, ts uint64) {
	e.MP -= ab.MPConsumption
	e.Cast = &Cast{AbilityID: ab.ID, TS: ts}
}

// Damage applies a direct damage component dd from entity source to entity e.
func (e *E) Damage(source *E, dd *ability.Damage) *ability.DamageFeedback {
	var amount uint64
	if dd.Amount >= e.HP {
		amount = e.HP
		e.HP = 0
		e.Dead = true
		e.Cast = nil
	} else {
		amount = dd.Amount
		e.HP -= dd.Amount
	}
	return &ability.DamageFeedback{Amount: amount}
}

// ApplyEffects applies all ability components.
func (e *E) ApplyEffects(source *E, effects []ability.Effect) ([]ability.EffectFeedback, error) {

	var result *multierror.Error
	var fbes []ability.EffectFeedback

	for _, effect := range effects {
		veffect := effect.GetValue()
		switch v := veffect.(type) {
		case *ability.Damage:
			fbes = append(fbes, ability.EffectFeedback{
				DamageFeedback: e.Damage(source, v),
			})
		case *ability.Heal:
			result = multierror.Append(result, gerrors.ErrNotImplementedYet{Version: "0.2.0"})
		case *ability.HealOverTime:
			result = multierror.Append(result, gerrors.ErrNotImplementedYet{Version: "0.2.0"})
		case *ability.DamageOverTime:
			result = multierror.Append(result, gerrors.ErrNotImplementedYet{Version: "0.2.0"})
		default:
			result = multierror.Append(result, gerrors.ErrNotImplementedYet{Version: "0.2.0"})
		}
	}

	return fbes, result.ErrorOrNil()
}

// ApplyEffectFeedbacks applies all feedback effects.
func (e *E) ApplyEffectFeedbacks(source *E, effects []ability.EffectFeedback) error {
	return nil // TODO
}
