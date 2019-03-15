package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/item"
)

func (h *handler) item(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.postItems(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *handler) postItems(w http.ResponseWriter, r *http.Request) {
	logger := log.With().Str("method", "POST").Str("route", "/item").Logger()

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var items []item.I
	if err := decoder.Decode(&items); err != nil {
		logger.Error().Err(err).Msg("invalid JSON")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger.Info().Int("items", len(items)).Msg("found")

	for _, it := range items {
		if err := h.ItemStore.SetItem(it); err != nil {
			logger.Error().Err(err).Str("item", it.ID.String()).Msg("failed to set item")
			http.Error(w, "store failure", http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
