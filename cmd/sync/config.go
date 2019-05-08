package main

import (
	"errors"

	"github.com/elojah/game_01/pkg/ulid"
)

// Config is the udp server structure config.
type Config struct {
	ID         ulid.ID `json:"id"`
	TickRate   uint32  `json:"tick_rate"`
	BatchSize  uint32  `json:"batch_size"`
	EntityPort uint    `json:"entity_port"`
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

	cTickRate, ok := fconf["tick_rate"]
	if !ok {
		return errors.New("missing key tick_rate")
	}
	cTickRateFloat64, ok := cTickRate.(float64)
	if !ok {
		return errors.New("key tick_rate invalid. must be numeric")
	}
	c.TickRate = uint32(cTickRateFloat64)

	cBatchSize, ok := fconf["batch_size"]
	if !ok {
		return errors.New("missing key batch_size")
	}
	cBatchSizeFloat64, ok := cBatchSize.(float64)
	if !ok {
		return errors.New("key batch_size invalid. must be numeric")
	}
	c.BatchSize = uint32(cBatchSizeFloat64)

	cEntityPort, ok := fconf["entity_port"]
	if !ok {
		return errors.New("missing key entity_port")
	}
	cEntityPortFloat64, ok := cEntityPort.(float64)
	if !ok {
		return errors.New("key entity_port invalid. must be numeric")
	}
	c.EntityPort = uint(cEntityPortFloat64)

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

	return nil
}
