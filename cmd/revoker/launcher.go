package main

import (
	"sync"

	"github.com/elojah/services"
)

// Namespaces maps configs used for auth server.
type Namespaces struct {
	Service services.Namespace
}

// Launcher represents service auth server launcher.
type Launcher struct {
	*services.Configs
	ns Namespaces

	service *service
	m       sync.Mutex
}

// NewLauncher returns service new auth server Launcher.
func (service *service) NewLauncher(ns Namespaces, nsRead ...services.Namespace) *Launcher {
	return &Launcher{
		Configs: services.NewConfigs(nsRead...),
		service: service,
		ns:      ns,
	}
}

// Up starts the auth server service with new configs.
func (l *Launcher) Up(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	sconfig := Config{}
	if err := sconfig.Dial(configs[l.ns.Service]); err != nil {
		return err
	}
	return l.service.Dial(sconfig)
}

// Down stops the auth server service.
func (l *Launcher) Down(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	return l.service.Close()
}
