package main

import (
	"errors"
)

// Config is web quic server structure config.
type Config struct {
	Lifespan uint64
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return c.Lifespan == rhs.Lifespan
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}

	cLifespan, ok := fconf["lifespan"]
	if !ok {
		return errors.New("missing key lifespan")
	}
	cLifespanFloat64, ok := cLifespan.(float64)
	if !ok {
		return errors.New("key lifespan invalid. must be numeric")
	}
	c.Lifespan = uint64(cLifespanFloat64)

	return nil
}
