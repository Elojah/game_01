package network

import (
	"sync"

	"github.com/elojah/services"
)

// Namespaces maps configs used for auth server.
type Namespaces struct {
	Client services.Namespace
}

// Launcher represents a auth server launcher.
type Launcher struct {
	*services.Configs
	ns Namespaces

	client *Client
	m      sync.Mutex
}

// NewLauncher returns a new auth server Launcher.
func (client *Client) NewLauncher(ns Namespaces, nsRead ...services.Namespace) *Launcher {
	return &Launcher{
		Configs: services.NewConfigs(nsRead...),
		client:  client,
		ns:      ns,
	}
}

// Up starts the auth server service with new configs.
func (l *Launcher) Up(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	sconfig := Config{}
	if err := sconfig.Dial(configs[l.ns.Client]); err != nil {
		return err
	}
	return l.client.Dial(sconfig)
}

// Down stops the auth server service.
func (l *Launcher) Down(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	return l.client.Close()
}
