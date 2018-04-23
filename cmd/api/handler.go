package main

import (
	"github.com/gocql/gocql"
	"github.com/sirupsen/logrus"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/dto"
	"github.com/elojah/game_01/storage"
	"github.com/elojah/nats-streaming"
	"github.com/elojah/udp"
)

type handler struct {
	*logrus.Entry
	game.Services
	Queue stan.Service
}

func (h handler) Route(mux *udp.Mux, cfg Config) {
	mux.Handler = h.handle
}

func (h *handler) handle(packet udp.Packet) error {

	// # Set logger.
	ip := packet.Source.String()
	logger := h.WithFields(logrus.Fields{
		"id":     packet.ID,
		"source": ip,
	})

	// # Unmarshal message.
	msg := dto.Message{}
	if _, err := msg.Unmarshal(packet.Data); err != nil {
		logger.WithFields(logrus.Fields{
			"status": "unmarshalable",
			"error":  err,
		}).Error("packet rejected")
		return err
	}

	// # Parse message UUID.
	uuid, err := gocql.UUIDFromBytes(msg.Token[:])
	if err != nil {
		logger.WithFields(logrus.Fields{
			"status": "unformatted",
			"error":  err,
		}).Error("packet rejected")
		return err
	}

	// # Search message UUID in storage.
	tokens, err := h.ListToken(game.TokenSubset{
		IDs: []game.ID{uuid},
	})
	if len(tokens) == 0 {
		err = storage.ErrNotFound
	}
	if err != nil {
		logger.WithFields(logrus.Fields{
			"status": "unidentified",
			"error":  err,
		}).Error("packet rejected")
		return err
	}
	token := tokens[0]

	// # Match message UUID with source IP.
	if token.IP.String() != ip {
		logger.WithFields(logrus.Fields{
			"status": "hijack",
			"error":  err,
		}).Error("packet rejected")
		return err
	}

	// TODO set last ack of current token/user in a ack service
	if msg.ACK != nil {
		go func() {
		}()
	}

	switch msg.Action.(type) {
	case dto.Attack:
		go func() { _ = h.attack(logger.WithField("action", "attack"), msg.Action.(dto.Attack)) }()
	case dto.Move:
		go func() { _ = h.move(logger.WithField("action", "move"), msg.Action.(dto.Move)) }()
	}

	return nil
}

func (h *handler) attack(logger *logrus.Entry, a dto.Attack) error {
	// TODO remove hp from actor to target with actor service scylla only
	return nil
}

func (h *handler) move(logger *logrus.Entry, m dto.Move) error {
	// TODO move player
	return nil
}
