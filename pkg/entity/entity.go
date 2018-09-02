package entity

import (
	"time"

	"github.com/elojah/game_01/pkg/ability"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/ulid"
)

// Store is an interface for E object.
type Store interface {
	SetEntity(E, int64) error
	GetEntity(Subset) (E, error)
	DelEntity(Subset) error
}

// Subset is a subset to retrieve one entity.
type Subset struct {
	ID     ulid.ID
	MinTS  int64
	MaxTS  int64
	Cursor uint64
	Count  int64
}

// Service represents entity usecases.
type Service interface {
	Disconnect(id ulid.ID, tok account.Token) error
}

// CastAbility decreases MP (without check) and assign a new cast to entity.
func (e *E) CastAbility(ab ability.A, ts time.Time) {
	e.MP -= ab.MPConsumption
	e.Cast = &Cast{AbilityID: ab.ID, TS: ts}
}

// Damage applies a direct damage component dd from entity source to entity e.
func (e *E) Damage(source E, dd ability.Damage) *ability.DamageFeedback {
	if dd.Amount >= e.HP {
		e.HP = 0
		e.Dead = true
	} else {
		e.HP -= dd.Amount
	}
	return &ability.DamageFeedback{Amount: dd.Amount}
}
