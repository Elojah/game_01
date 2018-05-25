package main

import (
	"errors"
)

// Config is the udp server structure config.
type Config struct {
	TickRate uint32 `json:"tick_rate"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return c == rhs
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}

	cTickRate, ok := fconf["tick_rate"]
	if !ok {
		return errors.New("missing key tick_rate")
	}
	cTickRateFloat64, ok := cTickRate.(float64)
	if !ok {
		return errors.New("key tick_rate invalid. must be numeric")
	}
	c.TickRate = uint32(cTickRateFloat64)

	return nil
}
