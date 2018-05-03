package main

import (
	"errors"
	"sync"

	"github.com/elojah/services"
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

// Namespaces maps configs used for api service with config file namespaces.
type Namespaces struct {
	Player services.Namespace
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

	return l.c.Dial(configs[l.ns.Player])
}

// Down stops the couchbase service.
func (l *Launcher) Down(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	return nil
}
