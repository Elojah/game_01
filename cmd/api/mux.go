package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/elojah/game_01"
)

// Handler is handle function responsible to process incoming data.
type Handler func([]byte) error

// Mux handles data and traffic parameters.
type Mux struct {
	*logrus.Entry
	*Config

	net.Conn
	sync.Map

	game.Services
}

// NewMux returns a new clear Mux.
func NewMux() *Mux {
	return &Mux{}
}

// Add adds a new handler identified by a string.
func (m *Mux) Add(identifier string, f Handler) {
	m.Store(identifier, f)
}

// Get returns a previously added handler identified by a string.
func (m *Mux) Get(identifier string) (Handler, error) {
	f, ok := m.Load(identifier)
	if !ok {
		return nil, fmt.Errorf("handler %s doesn't exist", identifier)
	}
	return f.(Handler), nil
}

// Read reads one 1024 packet from Conn and run it in identifier handler.
func (m *Mux) Read() error {
	for {
		raw := make([]byte, 1024)
		_, err := m.Conn.Read(raw)
		if err != nil {
			m.Logger.WithField("error", err).Error("failed to read")
			return err
		}
		m.Logger.WithField("data", string(raw)).Info("received packet")
	}
}
