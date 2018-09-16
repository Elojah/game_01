package main

import (
	"errors"
)

// Config is client config.
// Tolerance in milliseconds.
type Config struct {
	Address   string `json:"address"`
	Tolerance uint64 `json:"tolerance"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return (c.Address != rhs.Address &&
		c.Tolerance == rhs.Tolerance)
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	var err error
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}
	cAddress, ok := fconf["address"]
	if !ok {
		return errors.New("missing key address")
	}
	c.Address, ok = cAddress.(string)
	if !ok {
		return errors.New("key address invalid. must be string")
	}

	cTolerance, ok := fconf["tolerance"]
	if !ok {
		return errors.New("missing key tolerance")
	}
	cToleranceFloat, ok := cTolerance.(float64)
	if !ok {
		return errors.New("key tolerance invalid. must be numeric")
	}
	c.Tolerance = uint64(cToleranceFloat)

	return nil
}
