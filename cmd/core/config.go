package main

import (
	"errors"

	"github.com/oklog/ulid"

	"github.com/elojah/game_01"
)

// Config is the udp server structure config.
type Config struct {
	ID          game.ID   `json:"id"`
	Limit       int       `json:"limit"`
	Movelerance float64   `json:"move_tolerance"`
	Cores       []game.ID `json:"cores"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	if len(c.Cores) != len(rhs.Cores) {
		return false
	}
	for i := range c.Cores {
		if c.Cores[i].Compare(rhs.Cores[i]) != 0 {
			return false
		}
	}
	return c.ID.Compare(rhs.ID) == 0 &&
		c.Limit == rhs.Limit &&
		c.Movelerance == rhs.Movelerance
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

	cMovelerance, ok := fconf["move_tolerance"]
	if !ok {
		return errors.New("missing key move_tolerance")
	}
	c.Movelerance, ok = cMovelerance.(float64)
	if !ok {
		return errors.New("key move_tolerance invalid. must be numeric")
	}

	cCores, ok := fconf["cores"]
	if !ok {
		return errors.New("missing key cores")
	}
	cCoresSlice, ok := cCores.([]interface{})
	if !ok {
		return errors.New("key cores invalid. must be slice")
	}
	c.Cores = make([]game.ID, len(cCoresSlice))
	for i, core := range cCoresSlice {
		coreString, ok := core.(string)
		if !ok {
			return errors.New("value in cores invalid. must be string")
		}
		var err error
		c.Cores[i], err = ulid.Parse(coreString)
		if err != nil {
			return err
		}
	}
	return nil
}
