package main

import (
	"sync"

	"github.com/elojah/services"
)

// Namespaces maps configs used for api service with config file namespaces.
type Namespaces struct {
	Service services.Namespace
}

// Launcher represents svc api launcher.
type Launcher struct {
	*services.Configs
	ns Namespaces

	svc *service
	m   sync.Mutex
}

// NewLauncher returns svc new couchbase Launcher.
func (svc *service) NewLauncher(ns Namespaces, nsRead ...services.Namespace) *Launcher {
	return &Launcher{
		Configs: services.NewConfigs(nsRead...),
		svc:     svc,
		ns:      ns,
	}
}

// Up starts the couchbase service with new configs.
func (l *Launcher) Up(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	sconfig := Config{}
	if err := sconfig.Dial(configs[l.ns.Service]); err != nil {
		return err
	}
	return l.svc.Dial(sconfig)
}

// Down stops the couchbase service.
func (l *Launcher) Down(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	return nil
}
