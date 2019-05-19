package main

import (
	"sync"

	"github.com/elojah/services"
)

// Namespaces maps configs used for core service with config file namespaces.
type Namespaces struct {
	Service services.Namespace
}

// Launcher represents svc core launcher.
type Launcher struct {
	*services.Configs
	ns Namespaces

	svc *service
	m   sync.Mutex
}

// NewLauncher returns svc new config Launcher.
func (svc *service) NewLauncher(ns Namespaces, nsRead ...services.Namespace) *Launcher {
	return &Launcher{
		Configs: services.NewConfigs(nsRead...),
		svc:     svc,
		ns:      ns,
	}
}

// Up starts the config service with new configs.
func (l *Launcher) Up(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	sconfig := Config{}
	if err := sconfig.Dial(configs[l.ns.Service]); err != nil {
		return err
	}
	return l.svc.Dial(sconfig)
}

// Down stops the config service.
func (l *Launcher) Down(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	return l.svc.Close()
}
