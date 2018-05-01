package main

import (
	"errors"
	"sync"
	"time"

	"github.com/elojah/services"
)

// Config is the udp server structure config.
type Config struct {
	Tolerance time.Duration `json:"tolerance"`
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
	return true
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	var err error
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}
	cTolerance, ok := fconf["tolerance"]
	if !ok {
		return errors.New("missing key tolerance")
	}
	cToleranceString, ok := cTolerance.(string)
	if !ok {
		return errors.New("key tolerance invalid. must be string")
	}
	c.Tolerance, err = time.ParseDuration(cToleranceString)
	if err != nil {
		return err
	}
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
