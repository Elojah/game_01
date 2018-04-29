package main

import (
	"context"
	"time"

	"github.com/elojah/game_01/dto"
)

func (h *handler) attack(ctx context.Context, a dto.Attack, ts time.Time) error {
	// TODO remove hp from actor to target with actor service scylla only
	return nil
}
