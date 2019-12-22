package ui

import (
	"sync"

	"github.com/elojah/services"
)

// Namespaces maps configs used for ui.
type Namespaces struct {
	UI services.Namespace
}

// Launcher represents a auth server launcher.
type Launcher struct {
	*services.Configs
	ns Namespaces

	t *Term
	m sync.Mutex
}

// NewLauncher returns a new auth server Launcher.
func (t *Term) NewLauncher(ns Namespaces, nsRead ...services.Namespace) *Launcher {
	return &Launcher{
		Configs: services.NewConfigs(nsRead...),
		t:       t,
		ns:      ns,
	}
}

// Up starts the auth server service with new configs.
func (l *Launcher) Up(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	sconfig := Config{}
	if err := sconfig.Dial(configs[l.ns.UI]); err != nil {
		return err
	}
	return l.t.Dial(sconfig)
}

// Down stops the auth server service.
func (l *Launcher) Down(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	return l.t.Close()
}
