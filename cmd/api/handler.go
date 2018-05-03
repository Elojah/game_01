package main

import (
	"context"
	"net"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/dto"
	"github.com/elojah/mux"
)

type handler struct {
	Config
	game.Services
}

func (h handler) Route(m *mux.M, cfg Config) {
	m.Handler = h.handle
}

func (h *handler) handle(ctx context.Context, raw []byte) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	// # Unmarshal message.
	msg := dto.Message{}
	if _, err := msg.Unmarshal(raw); err != nil {
		logger.Error().Err(err).Str("status", "unmarshalable").Msg("packet rejected")
		return err
	}

	// # Parse message UUID.
	uuid := game.ID(msg.Token)

	// # Search message UUID in storage.
	token, err := h.GetToken(uuid)
	if err != nil {
		logger.Error().Err(err).Str("status", "unidentified").Str("uuid", uuid.String()).Msg("packet rejected")
		return err
	}

	// # Match message UUID with source IP.
	expected, _, _ := net.SplitHostPort(token.IP.String())
	actual, _, _ := net.SplitHostPort(ctx.Value(mux.Key("addr")).(string))
	if expected != actual {
		err := game.ErrWrongIP
		logger.Error().Err(err).Str("status", "hijacked").Str("expected", expected).Str("actual", actual).Msg("packet rejected")
		return err
	}

	// # Check TS in tolerance range.
	ts := time.Unix(0, msg.TS)
	now := time.Now()
	if ts.After(now) || now.Sub(ts) > h.Tolerance {
		err := game.ErrInvalidTS
		logger.Error().Err(err).Str("status", "timeout").Int64("ts", ts.UnixNano()).Int64("now", now.UnixNano()).Msg("packet rejected")
		return err
	}

	// TODO set last ack of current token/user in a ack service
	if msg.ACK != nil {
		go func() {
		}()
	}

	switch msg.Action.(type) {
	case dto.Attack:
		go func() { _ = h.attack(ctx, msg.Action.(dto.Attack), ts) }()
	case dto.Move:
		go func() { _ = h.move(ctx, msg.Action.(dto.Move), ts) }()
	}

	return nil
}
