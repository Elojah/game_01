package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/dto"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/game_01/pkg/usecase/token"
	"github.com/elojah/mux"
	"github.com/elojah/mux/client"
)

type handler struct {
	*mux.M
	*client.C

	event.QMapper

	token.T

	port      uint
	tolerance time.Duration
}

func (h *handler) Dial(c Config) error {
	h.M.Handler = h.handle
	h.port = c.ACKPort
	h.tolerance = c.Tolerance
	h.M.Listen()
	return nil
}

func (h *handler) Close() error {
	if err := h.M.Close(); err != nil {
		return err
	}
	return h.C.Close()
}

func (h *handler) handle(ctx context.Context, raw []byte) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	// #Unmarshal message.
	msg := dto.Event{}
	if _, err := msg.Unmarshal(raw); err != nil {
		logger.Error().Err(err).Str("status", "unmarshalable").Msg("packet rejected")
		return err
	}

	// #Parse message UUID.
	tokenID := ulid.ID(msg.Token)

	// #Get and check token.
	tok, err := h.T.Get(tokenID, ctx.Value(mux.Key("addr")).(string))
	if err != nil {
		logger.Error().Err(err).Str("status", "unidentified").Str("tokenID", ulid.String(tokenID)).Msg("failed to identify")
		return err
	}

	// #Send ACK to client.
	id := ulid.ID(msg.Token)
	ack := dto.ACK{ID: id}
	raw, err = ack.Marshal(nil)
	if err != nil {
		logger.Error().Err(err).Str("status", "internal").Msg("failed to marshal ack")
		return err
	}
	address := *tok.IP
	address.Port = int(h.port)
	go h.Send(raw, &address)

	// #Check TS in tolerance range.
	ts := time.Unix(0, msg.TS)
	now := time.Now()
	if ts.After(now) || now.Sub(ts) > h.tolerance {
		err := account.ErrInvalidTS
		logger.Error().Err(err).Str("status", "timeout").Int64("ts", ts.UnixNano()).Int64("now", now.UnixNano()).Msg("packet rejected")
		return err
	}

	// #Dispatch on actin t
	switch msg.Action.(type) {
	case dto.Move:
		go func() { _ = h.move(ctx, msg) }()
	case dto.Cast:
		go func() { _ = h.cast(ctx, msg) }()
	default:
		logger.Error().Msg("unrecognized action")
	}

	return nil
}
