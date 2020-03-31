package main

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/account/dto"
	gerrors "github.com/elojah/game_01/pkg/errors"
	"github.com/elojah/game_01/pkg/ulid"
)

func (h *handler) signin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := log.With().Str("route", "/signin").Str("method", "POST").Str("address", r.RemoteAddr).Logger()

	// #Read body
	var ac dto.SignInAccount
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

	// #Create token from account
	tok, err := h.account.CreateToken(account.A{
		Username: ac.Username,
		Password: ac.Password,
	}, ac.Address)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create token from account")
		http.Error(w, "failed to signin", http.StatusInternalServerError)
		return
	}

	logger = logger.With().Str("token", tok.ID.String()).Logger()

	// #Marshal token for response
	raw, err := json.Marshal(tok)
	if err != nil {
		logger.Error().Err(err).Msg("failed to marshal token")
		http.Error(w, "failed to signin", http.StatusInternalServerError)
		return
	}

	// #Write response
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(raw); err != nil {
		logger.Error().Err(err).Msg("failed to write response")
		return
	}

	logger.Info().Msg("signin success")
}

func (h *handler) signout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger := log.With().Str("route", "/signout").Str("address", r.RemoteAddr).Logger()

	// #Read body
	var ac dto.SignoutAccount
	if err := json.NewDecoder(r.Body).Decode(&ac); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger = logger.With().Str("username", ac.Username).Logger()

	// #Retrieve account by username.
	a, err := h.account.Fetch(ac.Username)
	if err != nil {
		switch errors.Cause(err).(type) {
		case gerrors.ErrNotFound:
			logger.Error().Err(err).Msg("invalid username")
			http.Error(w, "invalid username", http.StatusBadRequest)
			return
		}
		logger.Error().Err(err).Msg("failed to retrieve account")
		http.Error(w, "failed to signout", http.StatusInternalServerError)
		return
	}
	logger = logger.With().Str("account", a.ID.String()).Str("token", a.Token.String()).Logger()
	if a.Token.Compare(ac.Token) != 0 {
		logger.Error().Err(err).Msg("invalid token")
		http.Error(w, "invalid token", http.StatusBadRequest)
		return
	}

	// #Reset account token
	tokID := a.Token
	a.Token = ulid.Zero()
	if err := h.account.Upsert(a); err != nil {
		logger.Error().Err(err).Msg("failed to set account")
		http.Error(w, "failed to reset account token", http.StatusInternalServerError)
		return
	}

	// #Retrieve account token
	tok, err := h.account.FetchTokenFromAddr(tokID, r.RemoteAddr)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve token")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Disconnect token
	if err := h.account.DisconnectToken(tok.ID); err != nil {
		logger.Error().Err(err).Msg("failed to disconnect token")
		http.Error(w, "failed to disconnect token", http.StatusInternalServerError)
		return
	}

	// #Close token recurrer.
	if err := h.recurrer.Erase(tok.ID); err != nil {
		switch errors.Cause(err).(type) {
		case gerrors.ErrNotFound:
		default:
			logger.Error().Err(err).Msg("failed to remove recurrer")
			http.Error(w, "failed to remove recurrer", http.StatusInternalServerError)
			return
		}
	}

	// #Delete token
	if err := h.account.RemoveToken(tok.ID); err != nil {
		logger.Error().Err(err).Msg("failed to delete token")
		http.Error(w, "failed to delete token", http.StatusInternalServerError)
		return
	}

	// #Write response
	w.WriteHeader(http.StatusOK)

	logger.Info().Msg("signout success")
}
