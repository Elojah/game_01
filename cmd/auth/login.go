package main

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/storage"
)

func (h handler) login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := log.With().Str("route", "/login").Logger()

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	// Unmarshal payload
	var accountPayload game.Account
	if err = json.Unmarshal(b, &accountPayload); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	// Search account in redis
	account, err := h.GetAccount(game.AccountBuilder{
		Username: accountPayload.Username,
	})
	if err != nil && err != storage.ErrNotFound {
		logger.Error().Err(err).Msg("failed to get account")
		http.Error(w, "failed to retrieve account", http.StatusInternalServerError)
		return
	}
	if account.Password != accountPayload.Password {
		err = game.ErrWrongCredentials
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

	// Create a new token
	token := game.Token{
		ID:      game.NewULID(),
		Account: account.ID,
		IP:      ip,
	}
	if err := h.CreateToken(token); err != nil {
		logger.Error().Err(err).Msg("failed to create token")
		http.Error(w, "failed to create token", http.StatusInternalServerError)
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
