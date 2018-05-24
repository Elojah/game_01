package main

import ()

// Config is the udp server structure config.
type Config struct {
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return true
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	return nil
}
