package main

import (
	"errors"

	"github.com/oklog/ulid"

	"github.com/elojah/game_01"
)

// Config is the udp server structure config.
type Config struct {
	ID            game.ID   `json:"id"`
	Limit         int       `json:"limit"`
	MoveTolerance float64   `json:"move_tolerance"`
	Listeners     []game.ID `json:"listeners"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	if len(c.Listeners) != len(rhs.Listeners) {
		return false
	}
	for i := range c.Listeners {
		if c.Listeners[i].Compare(rhs.Listeners[i]) != 0 {
			return false
		}
	}
	return c.ID == rhs.ID &&
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

	cListeners, ok := fconf["listeners"]
	if !ok {
		return errors.New("missing key listeners")
	}
	cListenersSlice, ok := cListeners.([]interface{})
	if !ok {
		return errors.New("key listeners invalid. must be slice")
	}
	c.Listeners = make([]game.ID, len(cListenersSlice))
	for i, listener := range cListenersSlice {
		listenerString, ok := listener.(string)
		if !ok {
			return errors.New("value in listeners invalid. must be string")
		}
		var err error
		c.Listeners[i], err = ulid.Parse(listenerString)
		if err != nil {
			return err
		}
	}

	return nil
}
