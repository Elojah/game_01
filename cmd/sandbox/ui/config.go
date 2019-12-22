package ui

import "errors"

type Config struct {
	Tiles []interface{}
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	_, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}
	// cAddress, ok := fconf["address"]
	// if !ok {
	// 	return errors.New("missing key address")
	// }

	return nil
}
