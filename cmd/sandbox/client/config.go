package client

import (
	"errors"
)

// Config is client config.
// Tolerance in milliseconds.
type Config struct {
	Interval int `json:"interval"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return (c.Interval != rhs.Interval)
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}

	cInterval, ok := fconf["interval"]
	if !ok {
		return errors.New("missing key interval")
	}
	cIntervalFloat, ok := cInterval.(float64)
	if !ok {
		return errors.New("key interval invalid. must be numeric")
	}
	c.Interval = int(cIntervalFloat)

	return nil
}
