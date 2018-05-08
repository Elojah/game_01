package main

import (
	"errors"
	"sync"

	"github.com/elojah/services"
)

// Config is web quic server structure config.
type Config struct {
	Version   string   `json:"version"`
	Resources []string `json:"resources"`
	TickRate  uint     `json:"tick_rate"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	if len(c.Resources) != len(rhs.Resources) {
		return false
	}
	for i := range c.Resources {
		if c.Resources[i] != rhs.Resources[i] {
			return false
		}
	}
	if c.Version != rhs.Version {
		return false
	}
	if c.TickRate != rhs.TickRate {
		return false
	}
	return true
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}
	cVersion, ok := fconf["version"]
	if !ok {
		return errors.New("missing key version")
	}
	if c.Version, ok = cVersion.(string); !ok {
		return errors.New("key version invalid. must be string")
	}
	cResource, ok := fconf["resources"]
	if !ok {
		return errors.New("missing key resources")
	}
	cResources, ok := cResource.([]interface{})
	if !ok {
		return errors.New("key resources invalid. must be array")
	}
	c.Resources = make([]string, len(cResources))
	for i := range cResources {
		if c.Resources[i], ok = cResources[i].(string); !ok {
			return errors.New("key resources invalid. must be string array")
		}
	}
	cTickRate, ok := fconf["tick_rate"]
	if !ok {
		return errors.New("missing key tick_rate")
	}
	cTickRateFloat, ok := cTickRate.(float64)
	if !ok {
		return errors.New("key tick_rate invalid. must be numeric")
	}
	c.TickRate = uint(cTickRateFloat)
	return nil
}

// Namespaces maps configs used for api service with config file namespaces.
type Namespaces struct {
	API services.Namespace
}

// Launcher represents a api launcher.
type Launcher struct {
	*services.Configs
	ns Namespaces

	c *Config
	m sync.Mutex
}

// NewLauncher returns a new couchbase Launcher.
func (c *Config) NewLauncher(ns Namespaces, nsRead ...services.Namespace) *Launcher {
	return &Launcher{
		Configs: services.NewConfigs(nsRead...),
		c:       c,
		ns:      ns,
	}
}

// Up starts the couchbase service with new configs.
func (l *Launcher) Up(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	return l.c.Dial(configs[l.ns.API])
}

// Down stops the couchbase service.
func (l *Launcher) Down(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	return nil
}
