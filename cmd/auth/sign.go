package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/storage"
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

	logger := log.With().Str("route", "/signin").Str("method", "POST").Str("addr", r.RemoteAddr).Logger()

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

	// #Create token from account
	tok, err := h.T.New(a, r.RemoteAddr)
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

	logger := log.With().Str("route", "/signout").Str("addr", r.RemoteAddr).Logger()

	// #Read body
	var ac SignoutAccount
	if err := json.NewDecoder(r.Body).Decode(&ac); err != nil {
		logger.Error().Err(err).Msg("payload invalid")
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	logger = logger.With().Str("username", ac.Username).Logger()

	// #Retrieve account by username.
	a, err := h.AccountService.GetAccount(account.Subset{Username: ac.Username})
	if err != nil {
		if err == storage.ErrNotFound {
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
	a.Token = ulid.ID{}
	if err := h.AccountService.SetAccount(a); err != nil {
		logger.Error().Err(err).Msg("failed to set account")
		http.Error(w, "failed to reset account token", http.StatusInternalServerError)
		return
	}

	// #Retrieve account token
	tok, err := h.T.Get(tokID, r.RemoteAddr)
	if err != nil {
		logger.Error().Err(err).Msg("failed to retrieve token")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Disconnect token
	if err := h.T.Disconnect(tok.ID); err != nil {
		logger.Error().Err(err).Msg("failed to disconnect token")
		http.Error(w, "failed to disconnect token", http.StatusInternalServerError)
		return
	}

	// #Delete token
	if err := h.DelToken(account.TokenSubset{ID: tok.ID}); err != nil {
		logger.Error().Err(err).Msg("failed to delete token")
		http.Error(w, "failed to delete token", http.StatusInternalServerError)
		return
	}

	logger.Info().Msg("signout success")

	// #Write response
	w.WriteHeader(http.StatusOK)
}
