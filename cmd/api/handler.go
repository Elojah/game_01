package main

import (
	"context"
	"net"

	"github.com/oklog/ulid"
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

	event   event.App
	account account.App

	port      uint
	tolerance uint64
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
	var msg event.DTO
	if err := msg.Unmarshal(raw); err != nil {
		logger.Error().Err(err).Str("status", "unmarshalable").Msg("packet rejected")
		return err
	}

	// #Get and check token.
	tok, err := h.account.FetchTokenFromAddr(msg.Token, ctx.Value(mux.Key("addr")).(string))
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
	logger.Info().Str("id", ack.ID.String()).Str("addr", addr.String()).Msg("send ack")
	go h.Send(raw, addr)

	// #Check TS in tolerance range.
	now := ulid.Now()
	ts := msg.ID.Time()
	if ts > now || now-ts > h.tolerance {
		err := gerrors.ErrInvalidTS{MsgID: msg.ID.String(), TS: ts, Now: now}
		logger.Error().Err(err).Str("status", "timeout").Uint64("ts", ts).Uint64("now", now).Msg("packet rejected")
		return err
	}

	// #Dispatch on action.
	switch msg.Query.GetValue().(type) {
	case *event.Move:
		if err := h.move(ctx, msg); err != nil {
			logger.Error().Err(err).Str("event", "move").Msg("failed to send event")
		}
	case *event.Cast:
		if err := h.cast(ctx, msg); err != nil {
			logger.Error().Err(err).Str("event", "cast").Msg("failed to send event")
		}
	case *event.Loot:
		if err := h.loot(ctx, msg); err != nil {
			logger.Error().Err(err).Str("event", "loot").Msg("failed to send event")
		}
	case *event.Consume:
		if err := h.consume(ctx, msg); err != nil {
			logger.Error().Err(err).Str("event", "consume").Msg("failed to send event")
		}
	default:
		logger.Error().Msg("unrecognized action")
	}

	return nil
}
