package main

import (
	"encoding/json"
	"net/http"

	"github.com/elojah/game_01/pkg/event"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (h *handler) entityMove(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.postEntityMoves(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *handler) postEntityMoves(w http.ResponseWriter, r *http.Request) {

	logger := log.With().Str("method", "POST").Str("route", "/entity/move").Logger()

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var moves []event.MoveTarget
	if err := decoder.Decode(&moves); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger.Info().Int("entity_templates", len(moves)).Msg("found")

	for _, move := range moves {

		// #Add entity to new sector and remove from previous.
		// TODO look for previous position
		if err := h.AddEntityToSector(move.Source, move.Position.SectorID); err != nil {
			logger.Error().Err(errors.Wrapf(err, "add entity %s to sector %s", move.Source.String(), move.Position.SectorID.String())).Msg("failed to add entity to sector")
			return
		}
		if err := h.RemoveEntityFromSector(move.Source, move.Position.SectorID); err != nil {
			logger.Error().Err(errors.Wrapf(err, "remove entity %s from sector %s", move.Source.String(), s.ID.String())).Msg("failed to remove entity from sector")
			return
		}

		// #Move target
		target.Position = move.Position
	}

	// #Write new target state.
	return errors.Wrapf(h.EntityStore.SetEntity(target, ts), "set entity %s for ts %d", move.Source.String(), ts)
}
