package main

import (
	"errors"

	"github.com/elojah/game_01/pkg/ulid"
)

// Config is the udp server structure config.
type Config struct {
	ID            ulid.ID `json:"id"`
	Limit         int     `json:"limit"`
	MoveTolerance float64 `json:"move_tolerance"`
	LootRadius    float64 `json:"loot_radius"`
	ConsumeRadius float64 `json:"consume_radius"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return c.ID.Compare(rhs.ID) == 0 &&
		c.Limit == rhs.Limit &&
		c.MoveTolerance == rhs.MoveTolerance
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	var err error
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}

	cID, ok := fconf["id"]
	if !ok {
		return errors.New("missing key id")
	}
	cIDString, ok := cID.(string)
	if !ok {
		return errors.New("key id invalid. must be string")
	}
	if c.ID, err = ulid.Parse(cIDString); err != nil {
		return err
	}

	cLimit, ok := fconf["limit"]
	if !ok {
		return errors.New("missing key limit")
	}
	cLimitFloat, ok := cLimit.(float64)
	if !ok {
		return errors.New("key limit invalid. must be numeric")
	}
	c.Limit = int(cLimitFloat)

	cMoveTolerance, ok := fconf["move_tolerance"]
	if !ok {
		return errors.New("missing key move_tolerance")
	}
	c.MoveTolerance, ok = cMoveTolerance.(float64)
	if !ok {
		return errors.New("key move_tolerance invalid. must be numeric")
	}

	cLootRadius, ok := fconf["loot_radius"]
	if !ok {
		return errors.New("missing key loot_radius")
	}
	c.LootRadius, ok = cLootRadius.(float64)
	if !ok {
		return errors.New("key loot_radius invalid. must be numeric")
	}

	cConsumeRadius, ok := fconf["consume_radius"]
	if !ok {
		return errors.New("missing key consume_radius")
	}
	c.ConsumeRadius, ok = cConsumeRadius.(float64)
	if !ok {
		return errors.New("key consume_radius invalid. must be numeric")
	}

	return nil
}
