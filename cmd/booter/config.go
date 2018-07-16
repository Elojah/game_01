package main

import (
	"errors"
)

// Config is the udp server structure config.
type Config struct {
	HTMLFile string `json:"html_file"`
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

	cHTMLFile, ok := fconf["html_file"]
	if !ok {
		return errors.New("missing key html_file")
	}
	if c.HTMLFile, ok = cHTMLFile.(string); !ok {
		return errors.New("key html_file invalid. must be string")
	}

	return nil
}
