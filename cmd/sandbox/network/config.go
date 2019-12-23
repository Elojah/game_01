package network

import (
	"errors"
)

// Config is client config.
// ACKTolerance in milliseconds.
type Config struct {
	Address       string `json:"address"`
	ACKTolerance  uint64 `json:"ack_tolerance"`
	OmitTolerance uint64 `json:"omit_tolerance"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return (c.Address != rhs.Address &&
		c.ACKTolerance == rhs.ACKTolerance)
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
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

	cACKTolerance, ok := fconf["ack_tolerance"]
	if !ok {
		return errors.New("missing key ack_tolerance")
	}
	cACKToleranceFloat, ok := cACKTolerance.(float64)
	if !ok {
		return errors.New("key ack_tolerance invalid. must be numeric")
	}
	c.ACKTolerance = uint64(cACKToleranceFloat)

	cOmitTolerance, ok := fconf["omit_tolerance"]
	if !ok {
		return errors.New("missing key omit_tolerance")
	}
	cOmitToleranceFloat, ok := cOmitTolerance.(float64)
	if !ok {
		return errors.New("key omit_tolerance invalid. must be numeric")
	}
	c.OmitTolerance = uint64(cOmitToleranceFloat)

	return nil
}
