package main

import (
	"context"
	"net"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/dto"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/mux"
)

type handler struct {
	*mux.M

	event.QMapper

	account.TokenMapper

	tolerance time.Duration
}

func (h *handler) Dial(c Config) error {
	h.M.Handler = h.handle
	h.tolerance = c.Tolerance
	h.M.Listen()
	return nil
}

func (h *handler) Close() error {
	return h.M.Close()
}

func (h *handler) handle(ctx context.Context, raw []byte) error {

	logger := log.With().Str("packet", ctx.Value(mux.Key("packet")).(string)).Logger()

	// #Unmarshal message.
	msg := dto.Message{}
	if _, err := msg.Unmarshal(raw); err != nil {
		logger.Error().Err(err).Str("status", "unmarshalable").Msg("packet rejected")
		return err
	}

	// #Parse message UUID.
	tokenID := ulid.ID(msg.Token)

	// #Search message UUID in storage.
	token, err := h.GetToken(account.TokenSubset{ID: tokenID})
	if err != nil {
		logger.Error().Err(err).Str("status", "unidentified").Str("tokenID", tokenID.String()).Msg("packet rejected")
		return err
	}

	// #Match message UUID with source IP.
	expected, _, _ := net.SplitHostPort(token.IP.String())
	actual, _, _ := net.SplitHostPort(ctx.Value(mux.Key("addr")).(string))
	if expected != actual {
		err := account.ErrWrongIP
		logger.Error().Err(err).Str("status", "hijacked").Str("expected", expected).Str("actual", actual).Msg("packet rejected")
		return err
	}

	// #Send ACK to client.
	id := ulid.ID(msg.Token)
	ack := dto.ACK{ID: [16]byte(id)}
	raw, err = ack.Marshal(nil)
	if err != nil {
		logger.Error().Err(err).Str("status", "internal").Msg("failed to marshal ack")
		return err
	}
	go h.Send(raw, token.IP)

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
