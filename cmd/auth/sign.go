package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/dto"
)

func (h *handler) signin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := log.With().Str("route", "/signin").Str("addr", r.RemoteAddr).Logger()

	// #Read body
	var adto dto.Account
	if err := json.NewDecoder(r.Body).Decode(&adto); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}
	if err := adto.Check(); err != nil {
		logger.Error().Err(err).Msg("account invalid")
		http.Error(w, "account invalid", http.StatusBadRequest)
		return
	}
	a := adto.Domain()

	// #Create token from account
	tok, err := h.T.New(a, r.RemoteAddr)
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

	logger.Info().Msg("signin success")

	// #Write response
	w.WriteHeader(http.StatusOK)
	w.Write(raw)
}

func (h *handler) signout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
