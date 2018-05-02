package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/elojah/game_01"
)

type handler struct {
	game.Services

	srv *http.Server
}

// Dial starts the auth server.
func (h *handler) Dial(c Config) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/token", h.token)
	h.srv = &http.Server{
		Addr:    c.Address,
		Handler: mux,
	}
	go func() { _ = h.srv.ListenAndServeTLS(c.Cert, c.Key) }()
	return nil
}

// Close shutdowns the server listening.
func (h *handler) Close() error {
	return h.srv.Shutdown(context.Background())
}

func (h handler) token(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// continue
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	// Unmarshal
	var account game.Account
	if err = json.Unmarshal(b, &account); err != nil {
		http.Error(w, "payload invalid", http.StatusBadRequest)
		return
	}

	// Search account in redis
	account, err = h.GetAccount(game.AccountBuilder{
		Username: account.Username,
		Password: account.Password,
	})
	if err != nil {
		http.Error(w, "payload invalid", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
