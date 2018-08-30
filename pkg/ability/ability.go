package ability

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"

	gerrors "github.com/elojah/game_01/pkg/errors"
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

func (a A) Check(targets map[string]Targets) error {
	// #For all ability components.
	for cid, comp := range a.Components {

		// #Retrieve targets for this component.
		target, ok := targets[cid]
		if !ok {
			return errors.Wrapf(gerrors.ErrMissingTarget, "component %s for ability %s", cid, a.ID.String())
		}

		if len(target.Positions) != 0 {
			return gerrors.ErrNotImplementedYet
		}

		// #Check target numbers.
		if uint64(len(target.Entities)) > comp.NTargets {
			return errors.Wrapf(gerrors.ErrTooManyTargets, "component %s for ability %s with targets max %d and given %d", cid, a.ID.String(), len(target.Entities), comp.NTargets)
		}
	}

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
		Components        map[string]Component

		LastUsed int64
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
	a.Components = alias.Components
	var err error
	if a.CD, err = time.ParseDuration(alias.CD); err != nil {
		return err
	}
	if a.CastTime, err = time.ParseDuration(alias.CastTime); err != nil {
		return err
	}
	a.LastUsed = time.Unix(0, alias.LastUsed)
	return nil
}
