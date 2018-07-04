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

	logger := log.With().Str("route", "/login").Logger()

	// #Read body
	var accountPayload account.A
	if err := json.NewDecoder(r.Body).Decode(&accountPayload); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	// #Create token from account
	token, err := h.T.New(accountPayload, r.RemoteAddr)
	if err != nil {
		logger.Error().Err(err).Str("account", accountPayload.ID.String()).Msg("failed to create token from account")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	// #Set a new listener for this token
	listener, err := h.L.New(token.ID)
	if err != nil {
		logger.Error().Err(err).Str("token", token.ID.String()).Msg("failed to create token listener")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	// Marshal token for response
	raw, err := json.Marshal(token)
	if err != nil {
		logger.Error().Err(err).Msg("failed to marshal token")
		http.Error(w, "failed to connect", http.StatusInternalServerError)
		return
	}

	// Write response
	w.WriteHeader(http.StatusOK)
	w.Write(raw)
}
