package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/errors"
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

	logger := log.With().Str("route", "/subscribe").Str("method", "POST").Str("address", r.RemoteAddr).Logger()

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
	_, err := h.AccountStore.GetAccount(a.Username)
	if err != nil && err != errors.ErrNotFound {
		logger.Error().Err(err).Msg("failed to get account")
		http.Error(w, "failed to check account unicity", http.StatusInternalServerError)
		return
	}
	if err != errors.ErrNotFound {
		logger.Error().Err(err).Msg("account username found")
		http.Error(w, "username already exists", http.StatusUnauthorized)
		return
	}

	// #Set account in redis
	if err = h.AccountStore.SetAccount(a); err != nil {
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
	if _, err := w.Write(raw); err != nil {
		logger.Error().Err(err).Msg("failed to write response")
		return
	}

	logger.Info().Msg("subscribe success")
}

func (h *handler) unsubscribe(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := log.With().Str("route", "/unsubscribe").Str("method", "POST").Str("address", r.RemoteAddr).Logger()

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
	a, err := h.AccountStore.GetAccount(ac.Username)
	if err != nil {
		if err != errors.ErrNotFound {
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
	if !a.Token.IsZero() {
		logger = logger.With().Str("token", a.Token.String()).Logger()
		if _, err := h.TokenService.Access(a.Token, r.RemoteAddr); err != nil {
			logger.Error().Err(err).Msg("failed to retrieve token")
			http.Error(w, "failed to disconnect", http.StatusInternalServerError)
			return
		}
		if err := h.TokenService.Disconnect(a.Token); err != nil {
			logger.Error().Err(err).Msg("failed to disconnect token")
			http.Error(w, "failed to disconnect", http.StatusInternalServerError)
			return
		}
		// #Delete token
		if err := h.DelToken(a.Token); err != nil {
			logger.Error().Err(err).Msg("failed to delete token")
			http.Error(w, "failed to delete token", http.StatusInternalServerError)
			return
		}
	}

	// #Delete all associated PCs.
	pcs, err := h.ListPC(a.ID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to list pcs")
		http.Error(w, "failed to delete pcs", http.StatusInternalServerError)
		return
	}
	for _, pc := range pcs {
		if err := h.DelPC(a.ID, pc.ID); err != nil {
			logger.Error().Err(err).Str("pc", pc.ID.String()).Msg("failed to delete pc")
			http.Error(w, "failed to delete pcs", http.StatusInternalServerError)
			return
		}
	}

	// #Delete PC left number.
	if err := h.DelPCLeft(a.ID); err != nil {
		logger.Error().Err(err).Msg("failed to delete pc left")
		http.Error(w, "failed to delete pcs", http.StatusInternalServerError)
		return
	}

	// #Delete account.
	if err := h.AccountStore.DelAccount(a.Username); err != nil {
		logger.Error().Err(err).Msg("failed to delete account")
		http.Error(w, "failed to delete account", http.StatusInternalServerError)
		return
	}

	// #Write response
	w.WriteHeader(http.StatusOK)

	logger.Info().Msg("unsubscribe success")
}
