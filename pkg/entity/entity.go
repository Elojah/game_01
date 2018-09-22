package entity

import (
	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/ulid"
)

// Store is an interface for E object.
type Store interface {
	SetEntity(E, int64) error
	GetEntity(ulid.ID, int64) (E, error)
	DelEntity(ulid.ID) error
	DelEntityByTS(ulid.ID, int64) error
}

// Service represents entity usecases.
type Service interface {
	Disconnect(id ulid.ID, tok account.Token) error
}

// CastAbility decreases MP (without check) and assign a new cast to entity.
func (e *E) CastAbility(ab ability.A, ts int64) {
	e.MP -= ab.MPConsumption
	e.Cast = &Cast{AbilityID: ab.ID, TS: ts}
}

// Damage applies a direct damage component dd from entity source to entity e.
func (e *E) Damage(source E, dd ability.Damage) *ability.DamageFeedback {
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
