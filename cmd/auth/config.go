package main

import (
	"errors"
)

// Config is web quic server structure config.
type Config struct {
	Address   string   `json:"address"`
	Cert      string   `json:"cert"`
	Key       string   `json:"key"`
	Listeners []string `json:"listeners"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	if len(c.Listeners) != len(rhs.Listeners) {
		return false
	}
	for i := range c.Listeners {
		if c.Listeners[i] != rhs.Listeners[i] {
			return false
		}
	}
	return (c.Address == rhs.Address &&
		c.Cert == rhs.Cert &&
		c.Key == rhs.Key)
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
	if c.Address, ok = cAddress.(string); !ok {
		return errors.New("key address invalid. must be string")
	}

	cCert, ok := fconf["cert"]
	if !ok {
		return errors.New("missing key cert")
	}
	if c.Cert, ok = cCert.(string); !ok {
		return errors.New("key cert invalid. must be string")
	}

	cKey, ok := fconf["key"]
	if !ok {
		return errors.New("missing key key")
	}
	if c.Key, ok = cKey.(string); !ok {
		return errors.New("key key invalid. must be string")
	}

	cListeners, ok := fconf["listeners"]
	if !ok {
		return errors.New("missing key listeners")
	}
	cListenersSlice, ok := cListeners.([]interface{})
	if !ok {
		return errors.New("key listeners invalid. must be slice")
	}
	c.Listeners = make([]string, len(cListenersSlice))
	for i, listener := range cListenersSlice {
		c.Listeners[i], ok = listener.(string)
		if !ok {
			return errors.New("value in listeners invalid. must be string")
		}
	}

	return nil
}
