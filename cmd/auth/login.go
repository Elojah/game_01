package main

import (
	"encoding/json"
	"math/rand"
	"net"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/event"
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
	core := h.cores[rand.Intn(len(h.cores))]
	listener := event.Listener{ID: token.ID}
	if err := h.SendListener(listener, core); err != nil {
		logger.Error().Err(err).Msg("failed to create queue")
		http.Error(w, "failed to setup token", http.StatusInternalServerError)
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
