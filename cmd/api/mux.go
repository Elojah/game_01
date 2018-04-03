package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/elojah/game_01"
)

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

func (m *Mux) Add(identifier string, f Handler) {
	m.Store(identifier, f)
}

func (m *Mux) Get(identifier string) (Handler, error) {
	f, ok := m.Load(identifier)
	if !ok {
		return nil, fmt.Errorf("handler %s doesn't exist", identifier)
	}
	return f.(Handler), nil
}

func (m *Mux) Read() {
	for {
		raw := make([]byte, 1024)
		_, err := m.Conn.Read(raw)
		if err != nil {
			m.Logger.WithField("error", err).Error("failed to read")
			break
		}
	}
}
