package main

import (
	"errors"
	"time"
)

// Config is web quic server structure config.
type Config struct {
	Lifespan time.Duration
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

	var err error
	cLifespan, ok := fconf["lifespan"]
	if !ok {
		return errors.New("missing key lifespan")
	}
	cLifespanString, ok := cLifespan.(string)
	if !ok {
		return errors.New("key lifespan invalid. must be string")
	}
	if c.Lifespan, err = time.ParseDuration(cLifespanString); err != nil {
		return err
	}

	return nil
}
