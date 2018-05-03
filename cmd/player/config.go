package main

import (
	"errors"
)

// Config is the udp server structure config.
type Config struct {
	Subject string `json:"subject"`
	Bufsize int    `json:"bufsize"`
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

	cSubject, ok := fconf["subject"]
	if !ok {
		return errors.New("missing key subject")
	}
	c.Subject, ok = cSubject.(string)
	if !ok {
		return errors.New("key subject invalid. must be string")
	}

	cBufsize, ok := fconf["bufsize"]
	if !ok {
		return errors.New("missing key bufsize")
	}
	cBufsizeFloat64, ok := cBufsize.(float64)
	if !ok {
		return errors.New("key bufsize invalid. must be int")
	}
	c.Bufsize = int(cBufsizeFloat64)
	return nil
}
