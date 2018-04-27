package main

import (
	"context"
	"net"
	"time"

	"github.com/gocql/gocql"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/dto"
	"github.com/elojah/game_01/storage"
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

	// # Unmarshal message.
	msg := dto.Message{}
	if _, err := msg.Unmarshal(raw); err != nil {
		log.Ctx(ctx).Error().Err(err).Str("status", "unmarshalable").Msg("packet rejected")
		return err
	}

	// # Parse message UUID.
	uuid, err := gocql.UUIDFromBytes(msg.Token[:])
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Str("status", "unformatted").Msg("packet rejected")
		return err
	}

	// # Search message UUID in storage.
	tokens, err := h.ListToken(game.TokenSubset{
		IDs: []game.ID{uuid},
	})
	if err == nil && len(tokens) == 0 {
		err = storage.ErrNotFound
	}
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Str("status", "unidentified").Str("uuid", uuid.String()).Msg("packet rejected")
		return err
	}
	token := tokens[0]

	// # Match message UUID with source IP.
	expected, _, _ := net.SplitHostPort(token.IP.String())
	actual, _, _ := net.SplitHostPort(ctx.Value("address").(string))
	if expected != actual {
		err := game.ErrWrongIP
		log.Ctx(ctx).Error().Err(err).Str("status", "hijacked").Str("expected", expected).Str("actual", actual).Msg("packet rejected")
		return err
	}

	// # Check TS in tolerance range.
	ts := time.Unix(msg.TS, 0)
	now := time.Now()
	if ts.After(now) || now.Sub(ts) > h.Tolerance {
		err := game.ErrInvalidTS
		log.Ctx(ctx).Error().Err(err).Str("status", "hijacked").Int64("ts", ts.Unix()).Int64("now", now.Unix()).Msg("packet rejected")
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

func (h *handler) attack(ctx context.Context, a dto.Attack, ts time.Time) error {
	// TODO remove hp from actor to target with actor service scylla only
	return nil
}

func (h *handler) move(ctx context.Context, m dto.Move, ts time.Time) error {
	// h.Queue.
	// TODO move player
	return nil
}
