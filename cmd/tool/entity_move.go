package main

import (
	"encoding/json"
	"net/http"

	"github.com/elojah/game_01/pkg/event"
	gulid "github.com/elojah/game_01/pkg/ulid"
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

	var move event.Move
	if err := decoder.Decode(&move); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}
	logger.Info().Int("targets", len(move.Targets)).Msg("found")

	ts := gulid.NewID().Time()
	for _, targetID := range move.Targets {

		// #Get current entity state.
		e, err := h.EntityStore.GetEntity(targetID, ts)
		if err != nil {
			logger.Error().Err(errors.Wrapf(err, "get entity %s at ts %d", targetID.String(), ts)).Msg("failed to get entity")
			http.Error(w, "failed to retrieve entity", http.StatusInternalServerError)
			return
		}

		if e.Position.SectorID.Compare(move.Position.SectorID) != 0 {

			// #Add entity to new sector and remove from previous if necessary.
			if err := h.AddEntityToSector(targetID, move.Position.SectorID); err != nil {
				logger.Error().Err(errors.Wrapf(err, "add entity %s to sector %s", targetID.String(), move.Position.SectorID.String())).Msg("failed to add entity to sector")
				http.Error(w, "failed to add entity to new sector", http.StatusInternalServerError)
				return
			}
			if err := h.RemoveEntityFromSector(targetID, e.Position.SectorID); err != nil {
				logger.Error().Err(errors.Wrapf(err, "remove entity %s from sector %s", targetID.String(), e.Position.SectorID.String())).Msg("failed to remove entity from sector")
				http.Error(w, "failed to remove entity from previous sector", http.StatusInternalServerError)
				return
			}
		}

		// #Move target
		e.Position = move.Position

		// #Write new target state.
		if err := h.EntityStore.SetEntity(e, ts); err != nil {
			logger.Error().Err(errors.Wrapf(err, "set entity %s for ts %d", targetID.String(), ts))
			http.Error(w, "failed to set entity", http.StatusInternalServerError)
			return
		}
		logger.Info().Str("entity", e.ID.String()).Msg("tool move success")
	}

	w.WriteHeader(http.StatusOK)
}
