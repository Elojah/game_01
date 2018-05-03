package main

import (
	"errors"
	"time"
)

// Config is the udp server structure config.
type Config struct {
	Tolerance time.Duration `json:"tolerance"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return c == rhs
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	var err error
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}
	cTolerance, ok := fconf["tolerance"]
	if !ok {
		return errors.New("missing key tolerance")
	}
	cToleranceString, ok := cTolerance.(string)
	if !ok {
		return errors.New("key tolerance invalid. must be string")
	}
	c.Tolerance, err = time.ParseDuration(cToleranceString)
	if err != nil {
		return err
	}
	return nil
}
