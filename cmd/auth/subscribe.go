package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01"
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/storage"
)

func (h *handler) subscribe(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := log.With().Str("route", "/subscribe").Logger()

	// # Read body
	var a account.A
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}
	a.ID = game.NewID()

	// #Check username is unique
	_, err := h.AccountMapper.GetAccount(account.Subset{
		Username: a.Username,
	})
	if err != nil && err != storage.ErrNotFound {
		logger.Error().Err(err).Msg("failed to get account")
		http.Error(w, "failed to check account unicity", http.StatusInternalServerError)
		return
	}
	if err != storage.ErrNotFound {
		logger.Error().Err(err).Msg("account username found")
		http.Error(w, "username already exists", http.StatusUnauthorized)
		return
	}

	// #Set account in redis
	if err = h.AccountMapper.SetAccount(a); err != nil {
		logger.Error().Err(err).Msg("failed to create account")
		http.Error(w, "failed to create account", http.StatusInternalServerError)
		return
	}

	// #Add Permission to create X new chars.
	if err := h.SetPCLeft(entity.MaxPC, a.ID); err != nil {
		logger.Error().Err(err).Msg("failed to set character permission")
		http.Error(w, "failed to set permissions", http.StatusInternalServerError)
		return
	}

	// #Marshal token for response
	raw, err := json.Marshal(a)
	if err != nil {
		logger.Error().Err(err).Msg("failed to marshal account")
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		return
	}

	// #Write response
	w.WriteHeader(http.StatusOK)
	w.Write(raw)
}
