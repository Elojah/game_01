package main

import (
	"sync"

	"github.com/elojah/services"
)

// Namespaces maps configs used for api service with config file namespaces.
type Namespaces struct {
	Player services.Namespace
}

// Launcher represents a api launcher.
type Launcher struct {
	*services.Configs
	ns Namespaces

	a *app
	m sync.Mutex
}

// NewLauncher returns a new couchbase Launcher.
func (a *app) NewLauncher(ns Namespaces, nsRead ...services.Namespace) *Launcher {
	return &Launcher{
		Configs: services.NewConfigs(nsRead...),
		a:       a,
		ns:      ns,
	}
}

// Up starts the couchbase service with new configs.
func (l *Launcher) Up(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	sconfig := Config{}
	if err := sconfig.Dial(configs[l.ns.Player]); err != nil {
		// Add namespace key when returning error with logrus
		return err
	}
	return l.a.Dial(sconfig)
}

// Down stops the couchbase service.
func (l *Launcher) Down(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	return nil
}
