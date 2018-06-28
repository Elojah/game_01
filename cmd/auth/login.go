package main

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/event"
	"github.com/elojah/game_01/pkg/infra"
	"github.com/elojah/game_01/pkg/ulid"
	"github.com/elojah/game_01/storage"
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

	// # Read body
	var accountPayload account.A
	if err := json.NewDecoder(r.Body).Decode(&accountPayload); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	// Search account in redis
	a, err := h.AccountMapper.GetAccount(account.Subset{
		Username: accountPayload.Username,
	})
	if err != nil && err != storage.ErrNotFound {
		logger.Error().Err(err).Msg("failed to get account")
		http.Error(w, "failed to retrieve account", http.StatusInternalServerError)
		return
	}
	if a.Password != accountPayload.Password {
		err = account.ErrWrongCredentials
	}
	if err != nil {
		logger.Error().Err(err).Msg("failed to authenticate")
		http.Error(w, "wrong username/password", http.StatusUnauthorized)
		return
	}

	// Identify origin IP
	ip, err := net.ResolveUDPAddr("udp", r.RemoteAddr)
	if err != nil {
		logger.Error().Err(err).Str("address", r.RemoteAddr).Msg("failed to get valid IP")
		http.Error(w, "failed to identify ip", http.StatusInternalServerError)
		return
	}

	// Set a new token
	token := account.Token{
		ID:      ulid.NewID(),
		Account: a.ID,
		IP:      ip,
	}
	if err := h.SetToken(token); err != nil {
		logger.Error().Err(err).Msg("failed to create token")
		http.Error(w, "failed to create token", http.StatusInternalServerError)
		return
	}

	// Set a new listener for this token
	core, err := h.GetRandomCore(infra.CoreSubset{})
	if err != nil {
		if err == storage.ErrNotFound {
			logger.Error().Err(err).Msg("no core available")
			http.Error(w, "failed to create listener", http.StatusInternalServerError)
			return
		}
		logger.Error().Err(err).Msg("failed to get a core")
		http.Error(w, "failed to create listener", http.StatusInternalServerError)
		return
	}
	if err := h.SendListener(event.Listener{ID: token.ID, Action: event.Open}, core.ID); err != nil {
		logger.Error().Err(err).Str("core", core.ID.String()).Str("token", token.ID.String()).Msg("failed to add listener to token")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal token for response
	raw, err := json.Marshal(token)
	if err != nil {
		logger.Error().Err(err).Msg("failed to marshal token")
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		return
	}

	// Write response
	w.WriteHeader(http.StatusOK)
	w.Write(raw)
}
