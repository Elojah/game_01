package booter

import (
	"sync"

	"github.com/elojah/services"
)

// Namespaces maps configs used for auth server.
type Namespaces struct {
	Booter services.Namespace
}

// Launcher represents a auth server launcher.
type Launcher struct {
	*services.Configs
	ns Namespaces

	b *B
	m sync.Mutex
}

// NewLauncher returns a new auth server Launcher.
func (b *B) NewLauncher(ns Namespaces, nsRead ...services.Namespace) *Launcher {
	return &Launcher{
		Configs: services.NewConfigs(nsRead...),
		b:       b,
		ns:      ns,
	}
}

// Up starts the auth server service with new configs.
func (l *Launcher) Up(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	sconfig := Config{}
	if err := sconfig.Dial(configs[l.ns.Booter]); err != nil {
		return err
	}
	return l.b.Dial(sconfig)
}

// Down stops the auth server service.
func (l *Launcher) Down(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	return l.b.Close()
}
