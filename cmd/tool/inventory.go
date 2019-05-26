package main

import (
	"encoding/json"
	"net/http"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/rs/zerolog/log"
)

func (h *handler) inventoryHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.postInventories(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *handler) postInventories(w http.ResponseWriter, r *http.Request) {
	logger := log.With().Str("method", "POST").Str("route", "/inventory").Logger()

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var inventories []entity.Inventory
	if err := decoder.Decode(&inventories); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger.Info().Int("inventories", len(inventories)).Msg("found")

	for _, in := range inventories {
		if err := h.entity.UpsertInventory(in); err != nil {
			logger.Error().Err(err).Str("inventory", in.ID.String()).Msg("failed to set inventory")
			http.Error(w, "store failure", http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
