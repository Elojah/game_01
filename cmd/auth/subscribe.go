package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/storage"
	"github.com/elojah/game_01/pkg/ulid"
)

func (h *handler) subscribe(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := log.With().Str("route", "/subscribe").Str("method", "POST").Str("addr", r.RemoteAddr).Logger()

	// #Read body
	var ac Account
	if err := json.NewDecoder(r.Body).Decode(&ac); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}
	if err := ac.Check(); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}
	a := ac.Domain()
	a.ID = ulid.NewID()

	logger = logger.With().Str("account", a.ID.String()).Logger()

	// #Check username is unique
	_, err := h.AccountService.GetAccount(account.Subset{
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
	if err = h.AccountService.SetAccount(a); err != nil {
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

	logger.Info().Msg("subscribe success")

	// #Write response
	w.WriteHeader(http.StatusOK)
	w.Write(raw)
}

func (h *handler) unsubscribe(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := log.With().Str("route", "/unsubscribe").Str("method", "POST").Str("addr", r.RemoteAddr).Logger()

	// #Read body
	var ac Account
	if err := json.NewDecoder(r.Body).Decode(&ac); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}
	if err := ac.Check(); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	// #Search account in redis
	a, err := h.AccountService.GetAccount(account.Subset{
		Username: ac.Username,
	})
	if err != nil {
		if err != storage.ErrNotFound {
			logger.Error().Err(err).Msg("failed to get account")
			http.Error(w, "failed to retrieve account", http.StatusInternalServerError)
			return
		}
		logger.Error().Err(err).Msg("account invalid")
		http.Error(w, "wrong credentials", http.StatusBadRequest)
		return
	}
	logger = logger.With().Str("account", a.ID.String()).Logger()
	if a.Password != ac.Password {
		logger.Error().Err(err).Msg("account invalid")
		http.Error(w, "wrong credentials", http.StatusBadRequest)
		return
	}

	// #If account still connect, disconnect it.
	if !ulid.IsZero(a.Token) {
		logger = logger.With().Str("token", a.Token.String()).Logger()
		if _, err := h.T.Get(a.Token, r.RemoteAddr); err != nil {
			logger.Error().Err(err).Msg("failed to retrieve token")
			http.Error(w, "failed to disconnect", http.StatusInternalServerError)
			return
		}
		if err := h.T.Disconnect(a.Token); err != nil {
			logger.Error().Err(err).Msg("failed to disconnect token")
			http.Error(w, "failed to disconnect", http.StatusInternalServerError)
			return
		}
		// #Delete token
		if err := h.DelToken(account.TokenSubset{ID: a.Token}); err != nil {
			logger.Error().Err(err).Msg("failed to delete token")
			http.Error(w, "failed to delete token", http.StatusInternalServerError)
			return
		}
	}

	// #Delete all associated PCs.
	pcs, err := h.ListPC(entity.PCSubset{AccountID: a.ID})
	if err != nil {
		logger.Error().Err(err).Msg("failed to list pcs")
		http.Error(w, "failed to delete pcs", http.StatusInternalServerError)
		return
	}
	for _, pc := range pcs {
		if err := h.DelPC(entity.PCSubset{AccountID: a.ID, ID: pc.ID}); err != nil {
			logger.Error().Err(err).Str("pc", pc.ID.String()).Msg("failed to delete pc")
			http.Error(w, "failed to delete pcs", http.StatusInternalServerError)
			return
		}
	}

	// #Delete PC left number.
	if err := h.DelPCLeft(entity.PCLeftSubset{AccountID: a.ID}); err != nil {
		logger.Error().Err(err).Msg("failed to delete pc left")
		http.Error(w, "failed to delete pcs", http.StatusInternalServerError)
		return
	}

	// #Delete account.
	if err := h.AccountService.DelAccount(account.Subset{Username: a.Username}); err != nil {
		logger.Error().Err(err).Msg("failed to delete account")
		http.Error(w, "failed to delete account", http.StatusInternalServerError)
		return
	}

	// #Write response
	logger.Info().Msg("unsubscribe success")
	w.WriteHeader(http.StatusOK)
}
