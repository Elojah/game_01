package main

import (
	"context"
	"net/http"

	"github.com/elojah/game_01"
)

type handler struct {
	game.AccountMapper
	game.EntityMapper
	game.EntityTemplateMapper
	game.PCMapper
	game.PCLeftMapper
	game.QEventMapper
	game.QListenerMapper
	game.QRecurrerMapper
	game.SectorEntitiesMapper
	game.TokenMapper

	srv *http.Server

	cores []game.ID
	syncs []game.ID
}

// Dial starts the auth server.
func (h *handler) Dial(c Config) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", h.login)
	mux.HandleFunc("/subscribe", h.subscribe)
	mux.HandleFunc("/pc/create", h.createPC)
	mux.HandleFunc("/pc/list", h.listPC)
	mux.HandleFunc("/pc/conenct", h.connectPC)
	h.srv = &http.Server{
		Addr:    c.Address,
		Handler: mux,
	}
	go func() { _ = h.srv.ListenAndServeTLS(c.Cert, c.Key) }()
	h.cores = make([]game.ID, len(c.Cores))
	copy(h.cores, c.Cores)
	h.syncs = make([]game.ID, len(c.Syncs))
	copy(h.syncs, c.Syncs)
	return nil
}

// Close shutdowns the server listening.
func (h *handler) Close() error {
	return h.srv.Shutdown(context.Background())
}
