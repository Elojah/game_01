package main

import (
	"sync"

	"github.com/elojah/services"
)

// NamespacesHandler maps configs used for auth server.
type NamespacesHandler struct {
	Handler services.Namespace
}

// LauncherHandler represents a auth server launcher.
type LauncherHandler struct {
	*services.Configs
	ns NamespacesHandler

	h *handler
	m sync.Mutex
}

// NewLauncher returns a new auth server LauncherHandler.
func (h *handler) NewLauncher(ns NamespacesHandler, nsRead ...services.Namespace) *LauncherHandler {
	return &LauncherHandler{
		Configs: services.NewConfigs(nsRead...),
		h:       h,
		ns:      ns,
	}
}

// Up starts the auth server service with new configs.
func (l *LauncherHandler) Up(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	return l.h.Dial()
}

// Down stops the auth server service.
func (l *LauncherHandler) Down(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	return l.h.Close()
}
