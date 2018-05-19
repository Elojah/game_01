package main

import (
	"context"
	"net/http"

	"github.com/elojah/game_01"
)

type handler struct {
	game.AccountService
	game.PCLeftService
	game.QEventService
	game.QListenerService
	game.TokenService

	srv *http.Server

	cores []game.ID
}

// Dial starts the auth server.
func (h *handler) Dial(c Config) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", h.login)
	mux.HandleFunc("/subscribe", h.subscribe)
	h.srv = &http.Server{
		Addr:    c.Address,
		Handler: mux,
	}
	go func() { _ = h.srv.ListenAndServeTLS(c.Cert, c.Key) }()
	h.cores = make([]game.ID, len(c.Cores))
	copy(h.cores, c.Cores)
	return nil
}

// Close shutdowns the server listening.
func (h *handler) Close() error {
	return h.srv.Shutdown(context.Background())
}
