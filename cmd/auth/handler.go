package main

import (
	"context"
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
	mux.HandleFunc("/login", h.login)
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

func (h handler) login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// continue
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	account, err := h.GetAccount(game.AccountBuilder)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}
