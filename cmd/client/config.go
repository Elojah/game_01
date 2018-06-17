package main

import (
	"errors"
	"sync"

	game "github.com/elojah/game_01"
	"github.com/elojah/services"
	"github.com/oklog/ulid"
)

// Config is web quic server structure config.
type Config struct {
	Token    game.ID `json:"token"`
	TickRate uint    `json:"tick_rate"`
	Address  string  `json:"address"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	if c.Token.Compare(rhs.Token) != 0 {
		return false
	}
	if c.TickRate != rhs.TickRate {
		return false
	}
	if c.Address != rhs.Address {
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
	cTickRate, ok := fconf["tick_rate"]
	if !ok {
		return errors.New("missing key tick_rate")
	}
	cTickRateFloat, ok := cTickRate.(float64)
	if !ok {
		return errors.New("key tick_rate invalid. must be numeric")
	}
	c.TickRate = uint(cTickRateFloat)
	cToken, ok := fconf["token"]
	if !ok {
		return errors.New("missing key token")
	}
	cTokenString, ok := cToken.(string)
	if !ok {
		return errors.New("key token invalid. must be string")
	}
	var err error
	if c.Token, err = ulid.Parse(cTokenString); err != nil {
		return errors.New("key token invalid. must be ulid")
	}
	cAddress, ok := fconf["address"]
	if !ok {
		return errors.New("missing key address")
	}
	c.Address, ok = cAddress.(string)
	if !ok {
		return errors.New("key address invalid. must be string")
	}
	return nil
}

// Namespaces maps configs used for api service with config file namespaces.
type Namespaces struct {
	App services.Namespace
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

	return l.c.Dial(configs[l.ns.App])
}

// Down stops the couchbase service.
func (l *Launcher) Down(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	return nil
}
