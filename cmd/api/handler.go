package main

import (
	"github.com/sirupsen/logrus"

	"github.com/elojah/game_01/dto"
	"github.com/elojah/udp"
)

type handler struct {
	*logrus.Entry
}

func (h handler) Route(mux *udp.Mux, cfg Config) {
	mux.Handler = h.handle
}

func (h *handler) handle(packet udp.Packet) error {
	logger := h.WithField("id", packet.ID.String())

	msg := dto.Message{}
	if _, err := msg.Unmarshal(packet.Data); err != nil {
		logger.WithFields(logrus.Fields{
			"source": packet.Source.String(),
			"status": "unmarshalable",
			"error":  err,
		}).Error("packet rejected")
		return err
	}

	// TODO check ip/access token validity in an Token or IP service

	// TODO set last ack of current token/user in a ack service
	if msg.ACK != nil {
		go func() {
		}()
	}

	// TODO set position in tile38 + scylla actor service
	if msg.Position != nil {
		go func() {
		}()
	}

	switch msg.Action.(type) {
	case dto.Attack:
		return h.attack(logger.WithField("action", "attack"), msg.Action.(dto.Attack))
	}

	return nil
}

func (h *handler) attack(logger *logrus.Entry, a dto.Attack) error {
	// TODO remove hp from actor to target with actor service scylla only
	return nil
}
