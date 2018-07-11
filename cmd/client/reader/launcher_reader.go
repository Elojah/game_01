package main

import (
	"sync"

	"github.com/elojah/services"
)

// NamespacesReader maps configs used for auth server.
type NamespacesReader struct {
	Reader services.Namespace
}

// LauncherReader represents a auth server launcher.
type LauncherReader struct {
	*services.Configs
	ns NamespacesReader

	rd *reader
	m  sync.Mutex
}

// NewLauncher returns a new auth server LauncherReader.
func (rd *reader) NewLauncher(ns NamespacesReader, nsRead ...services.Namespace) *LauncherReader {
	return &LauncherReader{
		Configs: services.NewConfigs(nsRead...),
		rd:      rd,
		ns:      ns,
	}
}

// Up starts the auth server service with new configs.
func (l *LauncherReader) Up(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	sconfig := Config{}
	if err := sconfig.Dial(configs[l.ns.Reader]); err != nil {
		return err
	}
	return l.rd.Dial(sconfig)
}

// Down stops the auth server service.
func (l *LauncherReader) Down(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	return l.rd.Close()
}
