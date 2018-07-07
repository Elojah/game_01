package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
)

func (h *handler) login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := log.With().Str("route", "/login").Str("addr", r.RemoteAddr).Logger()

	// #Read body
	var accountPayload account.A
	if err := json.NewDecoder(r.Body).Decode(&accountPayload); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger = logger.With().Str("account", accountPayload.ID.String()).Logger()

	// #Create token from account
	tok, err := h.T.New(accountPayload, r.RemoteAddr)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create token from account")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	logger = logger.With().Str("token", tok.ID.String()).Logger()

	// #Set a new listener for this token
	listener, err := h.L.New(tok.ID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create token listener")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	logger = logger.With().Str("listener", listener.ID.String()).Logger()

	// #Marshal token for response
	raw, err := json.Marshal(tok)
	if err != nil {
		logger.Error().Err(err).Msg("failed to marshal token")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	logger.Info().Msg("login success")

	// #Write response
	w.WriteHeader(http.StatusOK)
	w.Write(raw)
}
