package main

import (
	"errors"
)

// Config is the udp server structure config.
// Tolerance in ms.
type Config struct {
	Tolerance int64 `json:"tolerance"`
	ACKPort   uint  `json:"ack_port"`
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
	cTolerance, ok := fconf["tolerance"]
	if !ok {
		return errors.New("missing key tolerance")
	}
	cToleranceFloat, ok := cTolerance.(float64)
	if !ok {
		return errors.New("key tolerance invalid. must be numeric")
	}
	c.Tolerance = int64(cToleranceFloat)

	cACKPort, ok := fconf["ack_port"]
	if !ok {
		return errors.New("missing key ack_port")
	}
	cACKPortFloat64, ok := cACKPort.(float64)
	if !ok {
		return errors.New("key ack_port invalid. must be numeric")
	}
	c.ACKPort = uint(cACKPortFloat64)

	return nil
}
