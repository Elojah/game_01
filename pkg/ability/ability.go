package ability

import (
	"encoding/json"
	"time"

	"github.com/elojah/game_01/pkg/ulid"
)

// Store is the communication interface for abilities.
type Store interface {
	SetAbility(A, ulid.ID) error
	GetAbility(Subset) (A, error)
	ListAbility(Subset) ([]A, error)
}

// Subset retrieves per EntityID+ID or list per EntityID.
type Subset struct {
	ID       ulid.ID
	EntityID ulid.ID
}

// UnmarshalJSON allows string as duration for cast time.
func (a *A) UnmarshalJSON(data []byte) error {
	var alias struct {
		ID                ulid.ID
		Type              ulid.ID
		Animation         ulid.ID
		Name              string
		MPConsumption     uint64
		PostMPConsumption uint64
		LastUsed          time.Time
		Components        map[string]Component

		CD       string
		CastTime string
	}
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	a.ID = alias.ID
	a.Type = alias.Type
	a.Animation = alias.Animation
	a.Name = alias.Name
	a.MPConsumption = alias.MPConsumption
	a.PostMPConsumption = alias.PostMPConsumption
	a.LastUsed = alias.LastUsed
	a.Components = alias.Components
	var err error
	if a.CD, err = time.ParseDuration(alias.CD); err != nil {
		return err
	}
	if a.CastTime, err = time.ParseDuration(alias.CastTime); err != nil {
		return err
	}
	return nil
}
