package entity

import (
	multierror "github.com/hashicorp/go-multierror"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/account"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/ulid"
)

// Store is an interface for E object.
type Store interface {
	SetEntity(E, uint64) error
	GetEntity(ulid.ID, uint64) (E, error)
	DelEntity(ulid.ID) error
	DelEntityByTS(ulid.ID, uint64) error
}

// Service represents entity usecases.
type Service interface {
	Disconnect(id ulid.ID, tok account.Token) error
}

// CastAbility decreases MP (without check) and assign a new cast to entity.
func (e *E) CastAbility(ab ability.A, ts uint64) {
	e.MP -= ab.MPConsumption
	e.Cast = &Cast{AbilityID: ab.ID, TS: ts}
}

// Damage applies a direct damage component dd from entity source to entity e.
func (e *E) Damage(source *E, dd ability.Damage) *ability.DamageFeedback {
	var amount uint64
	if dd.Amount >= e.HP {
		amount = e.HP
		e.HP = 0
		e.Dead = true
	} else {
		amount = dd.Amount
		e.HP -= dd.Amount
	}
	return &ability.DamageFeedback{Amount: amount}
}

// DamageFeedback applies a direct damage feedback dd from entity source to entity e.
func (e *E) DamageFeedback(source *E, dd ability.DamageFeedback) {
	if dd.Amount >= e.HP {
		e.HP = 0
		e.Dead = true
	} else {
		e.HP -= dd.Amount
	}
}

// ApplyEffects applies all ability components.
func (e *E) ApplyEffects(source *E, effects []ability.Effect) ([]ability.EffectFeedback, error) {

	var result *multierror.Error
	var fbes []ability.EffectFeedback

	for _, effect := range effects {
		veffect := effect.GetValue()
		switch v := veffect.(type) {
		case ability.Damage:
			fbes = append(fbes, ability.EffectFeedback{
				DamageFeedback: e.Damage(source, v),
			})
		case ability.Heal:
			result = multierror.Append(result, gerrors.ErrNotImplementedYet)
		case ability.HealOverTime:
			result = multierror.Append(result, gerrors.ErrNotImplementedYet)
		case ability.DamageOverTime:
			result = multierror.Append(result, gerrors.ErrNotImplementedYet)
		default:
			result = multierror.Append(result, gerrors.ErrNotImplementedYet)
		}
	}

	return fbes, result.ErrorOrNil()
}

// ApplyEffectFeedbacks applies all feedback effects.
func (e *E) ApplyEffectFeedbacks(source *E, effects []ability.EffectFeedback) error {

	var result *multierror.Error

	for _, effect := range effects {
		veffect := effect.GetValue()
		switch v := veffect.(type) {
		case ability.DamageFeedback:
			e.DamageFeedback(source, v)
		case ability.HealFeedback:
			result = multierror.Append(result, gerrors.ErrNotImplementedYet)
		case ability.HealOverTimeFeedback:
			result = multierror.Append(result, gerrors.ErrNotImplementedYet)
		case ability.DamageOverTimeFeedback:
			result = multierror.Append(result, gerrors.ErrNotImplementedYet)
		default:
			result = multierror.Append(result, gerrors.ErrNotImplementedYet)
		}
	}

	return result.ErrorOrNil()
}
