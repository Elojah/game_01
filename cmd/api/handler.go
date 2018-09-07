package main

import (
	"context"
	"net"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/infra"
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

	// #Get and check token.
	tok, err := h.TokenService.Access(msg.Token, ctx.Value(mux.Key("addr")).(string))
	if err != nil {
		logger.Error().Err(err).Str("status", "unidentified").Str("token", msg.Token.String()).Msg("failed to identify")
		return err
	}

	// #Send ACK to client.
	ack := infra.ACK{ID: msg.ID}
	raw, err = ack.Marshal()
	if err != nil {
		logger.Error().Err(err).Str("status", "internal").Msg("failed to marshal ack")
		return err
	}
	addr, err := net.ResolveUDPAddr("udp", tok.IP)
	if err != nil {
		logger.Error().Err(err).Msg("failed to parse ip")
		return err
	}
	addr.Port = int(h.port)
	go h.Send(raw, addr)

	// #Check TS in tolerance range.
	now := time.Now()
	if msg.TS.After(now) || now.Sub(msg.TS) > h.tolerance {
		err := gerrors.ErrInvalidTS
		logger.Error().Err(err).Str("status", "timeout").Int64("ts", msg.TS.UnixNano()).Int64("now", now.UnixNano()).Msg("packet rejected")
		return err
	}

	// #Dispatch on action.
	switch msg.Query.GetValue().(type) {
	case *event.Move:
		go func() { _ = h.move(ctx, msg) }()
	case *event.Cast:
		go func() { _ = h.cast(ctx, msg) }()
	default:
		logger.Error().Msg("unrecognized action")
	}

	return nil
}
