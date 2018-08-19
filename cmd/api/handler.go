package main

import (
	"context"
	"net"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/mux"
	"github.com/elojah/mux/client"
)

type handler struct {
	*mux.M
	*client.C

	event.QStore
	account.TokenService

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
	msg := event.DTO{}
	if err := msg.Unmarshal(raw); err != nil {
		logger.Error().Err(err).Str("status", "unmarshalable").Msg("packet rejected")
		return err
	}

	// #Parse message UUID.
	tokenID := ulid.ID(msg.Token)

	// #Get and check token.
	tok, err := h.TokenService.Access(tokenID, ctx.Value(mux.Key("addr")).(string))
	if err != nil {
		logger.Error().Err(err).Str("status", "unidentified").Str("tokenID", tokenID.String()).Msg("failed to identify")
		return err
	}

	// #Send ACK to client.
	id := ulid.ID(msg.Token)
	ack := infra.ACK{ID: id}
	raw, err = ack.Marshal()
	if err != nil {
		logger.Error().Err(err).Str("status", "internal").Msg("failed to marshal ack")
		return err
	}
	address, err := net.ResolveUDPAddr("udp", tok.IP)
	if err != nil {
		logger.Error().Err(err).Msg("failed to parse ip")
		return err
	}
	address.Port = int(h.port)
	go h.Send(raw, address)

	// #Check TS in tolerance range.
	ts := time.Unix(0, msg.TS)
	now := time.Now()
	if ts.After(now) || now.Sub(ts) > h.tolerance {
		err := account.ErrInvalidTS
		logger.Error().Err(err).Str("status", "timeout").Int64("ts", ts.UnixNano()).Int64("now", now.UnixNano()).Msg("packet rejected")
		return err
	}

	// #Dispatch on action.
	switch msg.Action.GetValue().(type) {
	case event.Move:
		go func() { _ = h.move(ctx, msg) }()
	case event.Cast:
		go func() { _ = h.cast(ctx, msg) }()
	default:
		logger.Error().Msg("unrecognized action")
	}

	return nil
}
