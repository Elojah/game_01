package main

import (
	"errors"

	"github.com/oklog/ulid"

	"github.com/elojah/game_01"
)

// Config is web quic server structure config.
type Config struct {
	Address string    `json:"address"`
	Cert    string    `json:"cert"`
	Key     string    `json:"key"`
	Cores   []game.ID `json:"cores"`
	Syncs   []game.ID `json:"syncs"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	if len(c.Cores) != len(rhs.Cores) {
		return false
	}
	for i := range c.Cores {
		if c.Cores[i].Compare(rhs.Cores[i]) != 0 {
			return false
		}
	}
	if len(c.Syncs) != len(rhs.Syncs) {
		return false
	}
	for i := range c.Syncs {
		if c.Syncs[i].Compare(rhs.Syncs[i]) != 0 {
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

	cCores, ok := fconf["cores"]
	if !ok {
		return errors.New("missing key cores")
	}
	cCoresSlice, ok := cCores.([]interface{})
	if !ok {
		return errors.New("key cores invalid. must be slice")
	}
	c.Cores = make([]game.ID, len(cCoresSlice))
	for i, core := range cCoresSlice {
		coreString, ok := core.(string)
		if !ok {
			return errors.New("value in cores invalid. must be string")
		}
		var err error
		c.Cores[i], err = ulid.Parse(coreString)
		if err != nil {
			return err
		}
	}

	cSyncs, ok := fconf["syncs"]
	if !ok {
		return errors.New("missing key syncs")
	}
	cSyncsSlice, ok := cSyncs.([]interface{})
	if !ok {
		return errors.New("key syncs invalid. must be slice")
	}
	c.Syncs = make([]game.ID, len(cSyncsSlice))
	for i, sync := range cSyncsSlice {
		syncString, ok := sync.(string)
		if !ok {
			return errors.New("value in syncs invalid. must be string")
		}
		var err error
		c.Syncs[i], err = ulid.Parse(syncString)
		if err != nil {
			return err
		}
	}

	return nil
}
