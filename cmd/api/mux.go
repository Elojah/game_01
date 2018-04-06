package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/cloudflare/golz4"
	"github.com/sirupsen/logrus"
)

// Handler is handle function responsible to process incoming data.
type Handler func([]byte) error

// Mux handles data and traffic parameters.
type Mux struct {
	*logrus.Entry
	*Config

	net.Conn
	sync.Map
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
		go func(raw []byte) {
			m.Logger.WithFields(logrus.Fields{
				"type":   "packet",
				"status": "received",
				"data":   string(raw),
			}).Info("")

			fbs := make([]byte, 1024)
			if err := lz4.Uncompress(raw, fbs); err != nil {
				m.Logger.WithFields(logrus.Fields{
					"type":   "packet",
					"format": "lz4",
					"status": "received",
					"error":  err,
				}).Info("")
				m.Logger.WithField("error", err).WithField("format", "lz4").Info("packet")
				return
			}
		}(raw)
	}
}
